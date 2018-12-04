#!/usr/bin/env escript

main(_) ->
	SupervisorPid = spawn(fun() -> startSupervisor() end),
	register(supervisor, SupervisorPid),
	parseLine(io:get_line(""), null, []),
	waitForEnd().

waitForEnd() ->
	receive
		quit -> ok
	end.


startSupervisor() ->
	runSupervisor(null, null, 0).


runSupervisor(MostSleepingGuard, MostSleptMinute, MaxDuration) ->
	receive
		{time_slept, GuardID, SleepingMinutes} ->
			% printHour(GuardID, SleepingMinutes),
			Duration = calculateGuardSleepDuration(SleepingMinutes, 0),
			{MaxSleptMinute, SleepsCount} = calculateMaxSleptMinute(SleepingMinutes, 0, 0),
			io:format(
				"#~s slept ~w (Part 1) minutes, ~w being the most frequent with ~w (Part 2) sleeps (Result: ~w)~n",
				[GuardID, Duration, MaxSleptMinute, SleepsCount, list_to_integer(binary_to_list(GuardID)) * MaxSleptMinute]
			 ),
			runSupervisor(MostSleepingGuard, MostSleptMinute, MaxDuration)
	end.


calculateGuardSleepDuration([], Duration) -> Duration;
calculateGuardSleepDuration([{_, Slept} | SleepingMinutes], Duration) when Slept > 0 ->
	calculateGuardSleepDuration(SleepingMinutes, Duration + 1);
calculateGuardSleepDuration([_ | SleepingMinutes], Duration) ->
	calculateGuardSleepDuration(SleepingMinutes, Duration).


calculateMaxSleptMinute([], MaxMinute, MaxSleeps) -> {MaxMinute, MaxSleeps};
calculateMaxSleptMinute([{Minute, Slept} | SleepingMinutes], _, MaxSleeps) when Slept > MaxSleeps ->
	calculateMaxSleptMinute(SleepingMinutes, Minute, Slept);
calculateMaxSleptMinute([_ | SleepingMinutes], MaxMinute, MaxSleeps) ->
	calculateMaxSleptMinute(SleepingMinutes, MaxMinute, MaxSleeps).


parseLine(eof, _, Guards) ->
	finishGuards(Guards);
parseLine(Line, GuardID, Guards) ->
	{NewGuardId, GuardsUpdated} = processLine(list_to_binary(Line), GuardID, Guards),
	parseLine(io:get_line(""), NewGuardId, GuardsUpdated).


finishGuards([]) -> ok;
finishGuards([Guard, Guards]) ->
	Guard ! finish,
	finishGuards(Guards).


processLine(Line, GuardID, Guards) ->
	% First char of the sentence defines the action
	Action = binary:part(Line, 19, 1),
	<<
		_:4/binary, "-", _:2/binary, "-", _:2/binary,
		" ", _:2/binary, ":", Minute:2/binary
	>> = binary:part(Line, 1, 16),
	action(
		GuardID,
		Line,
		list_to_integer(binary_to_list(Minute)),
		Action,
		Guards
	).


% Line indicating a guard starts a shift
action(_, Line, _, <<"G">>, Guards) ->
	% 40 == length before the id + length after (" begins shift\n" == 14)
	GuardID = binary:part(Line, 26, byte_size(Line) - 40),
	GuardName = binary_to_atom(GuardID, unicode),
	GuardsUpdated = getGuard(whereis(GuardName), GuardName, GuardID, Guards),
	{GuardID, GuardsUpdated};
% Line indicating a guard falls asleep
action(GuardID, _, Minute, <<"f">>, Guards) ->
	GuardName = binary_to_atom(GuardID, unicode),
	GuardName ! {sleep, Minute},
	{GuardID, Guards};
% Line indicating a guard wakes up
action(GuardID, _, Minute, <<"w">>, Guards) ->
	GuardName = binary_to_atom(GuardID, unicode),
	GuardName ! {wake, Minute},
	{GuardID, Guards};
action(_, _, _, Action, _) ->
	io:write("Invalid action: "),
	io:write(Action),
	io:write("~n"),
	null.


getGuard(undefined, GuardName, GuardID, Guards) ->
	Pid = spawn(fun() -> startGuard(GuardID) end),
	true = register(GuardName, Pid),
	[Pid, Guards];
getGuard(_, _, _, Guards) -> Guards.


startGuard(GuardID) ->
	processGuard(
		GuardID,
		[
			{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0},
			{9, 0}, {10, 0}, {11, 0}, {12, 0}, {13, 0}, {14, 0}, {15, 0}, {16, 0}, {17, 0},
			{18, 0}, {19, 0}, {20, 0}, {21, 0}, {22, 0}, {23, 0}, {24, 0}, {25, 0}, {26, 0},
			{27, 0}, {28, 0}, {29, 0}, {30, 0}, {31, 0}, {32, 0}, {33, 0}, {34, 0}, {35, 0},
			{36, 0}, {37, 0}, {38, 0}, {39, 0}, {40, 0}, {41, 0}, {42, 0}, {43, 0}, {44, 0},
			{45, 0}, {46, 0}, {47, 0}, {48, 0}, {49, 0}, {50, 0}, {51, 0}, {52, 0}, {53, 0},
			{54, 0}, {55, 0}, {56, 0}, {57, 0}, {58, 0}, {59, 0}
		],
		null
	).


processGuard(GuardID, MinutesSlept, MinuteStart) ->
	receive
		finish ->
			supervisor ! {time_slept, GuardID, MinutesSlept};
		{sleep, NewMinuteStart} ->
			processGuard(GuardID, MinutesSlept, NewMinuteStart);
		{wake, MinuteEnd} ->
			MinutesSlept1 = updateSlept(MinutesSlept, MinuteStart, MinuteEnd),
			processGuard(GuardID, MinutesSlept1, null)
	end.


updateSlept(Slept, Start, End) ->
	updateSleptTime(Slept, [], Start, End).


updateSleptTime([], SleptUpdated, _, _) ->
	lists:reverse(SleptUpdated);
updateSleptTime([{Minute, MinuteSlept} | Slept], SleptUpdated, Start, End) when Minute >= Start, Minute < End ->
	updateSleptTime(
		Slept,
		[{Minute, MinuteSlept + 1} | SleptUpdated],
		Start,
		End
	);
updateSleptTime([{Minute, MinuteSlept} | Slept], SleptUpdated, Start, End) ->
	updateSleptTime(
		Slept,
		[{Minute, MinuteSlept} | SleptUpdated],
		Start,
		End
	).

printHour(_, []) ->
	ok;
printHour(GuardID, [Head | Tail]) ->
	io:format("~s ~w~n", [GuardID, Head]),
	printHour(GuardID, Tail).
