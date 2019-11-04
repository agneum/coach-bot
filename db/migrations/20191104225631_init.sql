-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS autopost_plans(
id integer not null primary key AUTOINCREMENT,
chatid BIGINT NOT NULL,
type VARCHAR(16) NOT NULL,
texttemplate TEXT DEFAULT '',
lastscheduled TIMESTAMP,
intervals TEXT NOT NULL,
startdate TIMESTAMP NOT NULL,
enddate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
isactive integer NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS scheduled_posts (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	chatid BIGINT NOT NULL,
	senddate TIMESTAMP NOT NULL,
	message TEXT NOT NULL,
	done INTEGER NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE autopost_plans;
DROP TABLE scheduled_posts;
-- +goose StatementEnd
