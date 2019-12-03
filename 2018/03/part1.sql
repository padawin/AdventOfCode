-- Overlapping
SELECT
	DISTINCT coords,
	COUNT(coords)
FROM (
	-- All used
	SELECT
		id_input,
		GENERATE_SERIES(tlx, brx) || ',' || y AS coords
	FROM (
		-- Y used
		SELECT
			id_input,
			GENERATE_SERIES(tly, bry) AS y,
			tlx,
			brx
		FROM (
			-- Coordinates expansion
			SELECT
				id_input,
				x AS tlx,
				y AS tly,
				w,
				h,
				x+w-1 AS brx,
				y+h-1 AS bry
			FROM
				rectangle
		) AS rectangles_limits
	) AS rectangles_y
) AS coords_data
GROUP BY coords
HAVING COUNT(coords) > 1;
