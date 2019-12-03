-- Overlapping
WITH coords_data AS (
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
)
SELECT
	id_input,
	count(coords),
	sum(count_coord)
FROM (
	SELECT
		id_input,
		coords_data.coords,
		count_coord
	FROM
		coords_data
		JOIN (
			SELECT
				coords, count(coords) as count_coord
			FROM
				coords_data
			GROUP BY coords
		) AS coords_count USING (coords)
) AS aggreg
GROUP BY id_input
HAVING count(coords) = sum(count_coord)
