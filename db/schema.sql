CREATE TYPE match_result AS ENUM ('home', 'draw', 'away');

CREATE TABLE if not exists fixtures (
    id int PRIMARY KEY NOT NULL, 
    time timestamp with time zone,
    home_team_id int NOT NULL,
    away_team_id int NOT NULL,
    home_team_goals int,
    away_team_goals int,
    result match_result
);