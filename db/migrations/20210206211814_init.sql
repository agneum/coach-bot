-- +goose Up
-- SQL in this section is executed when the migration is applied.

create table coach
(
  id         int primary key not null,
  name       text            not null,
  created_at timestamp       not null,
  updated_at timestamp       not null
);

insert into coach(id, name, created_at, updated_at)
values
  (1, 'Павел', datetime('now'), datetime('now')),
  (2, 'Брюс', datetime('now'), datetime('now'));

create table place
(
  id          integer primary key not null,
  name        text unique         not null,
  address     text                not null default '',
  description text                not null default '',
  created_at  timestamp           not null,
  updated_at  timestamp           not null
);

insert into place(id, name, address, description, created_at, updated_at)
values
  (1, 'Зал Горностай', 'ул. Полевая, 5a', '', datetime('now'), datetime('now')),
  (2, 'Зал Юность', 'ул. Строителей, 23', '', datetime('now'), datetime('now')),
  (3, 'Зал Школы №121', 'ул. Труженников, 10', '', datetime('now'), datetime('now')),
  (4, 'СК Мой пляж', 'ул. Большая, 254/1', '', datetime('now'), datetime('now'));


create table template
(
  id          integer primary key                                 not null,
  title       text                                                not null,
  description text                                                not null default '',
  type        text check ( type in ('classic', 'beach', 'party')) not null,
  note        text                                                not null default '',
  coach_id    int,
  place_id    int                                                 not null,
  weekday     int check ( weekday >= 0 and weekday <= 6 )         not null,
  start_time  text                                                not null,
  duration    text                                                not null,
  created_at  timestamp                                           not null,
  updated_at  timestamp                                           not null,
  foreign key (coach_id) references coach (id),
  foreign key (place_id) references place (id)
);

insert into template(id, title, description, type, note, coach_id, place_id, weekday, start_time, duration, created_at,
                     updated_at)
values
  (1, '', '', 'classic', '', 1, 1, 1, '20:00', '2h', datetime('now'), datetime('now')),
  (2, '', '', 'classic', '', 1, 2, 2, '08:00', '2h', datetime('now'), datetime('now')),
  (3, '', '', 'classic', '', 1, 3, 2, '20:00', '2h', datetime('now'), datetime('now')),
  (4, '', '', 'beach', '', 2, 4, 3, '19:00', '1.5h', datetime('now'), datetime('now')),
  (5, '', '', 'classic', 'женская сетка', 1, 1, 3, '20:00', '2h', datetime('now'), datetime('now'))
;

create table event
(
  id          integer primary key                                 not null,
  title       text                                                not null,
  description text                                                not null default '',
  type        text check ( type in ('classic', 'beach', 'party')) not null,
  note        text                                                not null default '',
  coach_id    int,
  place_id    int                                                 not null,
  start_date  timestamp                                           not null,
  duration    text                                                not null,
  created_at  timestamp                                           not null,
  updated_at  timestamp                                           not null,
  foreign key (coach_id) references coach (id),
  foreign key (place_id) references place (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table event;
drop table template;
drop table place;
drop table coach;
