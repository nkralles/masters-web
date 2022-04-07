GRANT
ALL
PRIVILEGES
ON
DATABASE
masters TO masters;
ALTER
DATABASE masters SET intervalstyle = 'iso_8601';

create
extension if not exists timescaledb;
create
extension if not exists citext;

create table golfers
(
    player_id  int primary key not null,
    rank       int             not null,
    first_name citext          not null,
    last_name  citext          not null,
    cc         citext          not null
);

create index on golfers (rank);
create index on golfers (cc);
create index on golfers ((first_name || ' ' || last_name));

create table entries
(
    name          citext primary key not null,
    winning_score int                not null default 1000
);

create table user_golfer_entries
(
    uuid      uuid not null primary key default gen_random_uuid(),
    entry_name  citext references entries (name) on update cascade on delete cascade,
    golfer_id int references golfers (player_id) on update cascade on delete cascade
);

create table masters_holes
(
    hole_number int primary key not null,
    name        text            not null,
    par         int             not null,
    yards       int             not null
);
create index on masters_holes (hole_number);

insert into masters_holes(hole_number, name, par, yards)
VALUES (1, 'Tea Ollive', 4, 445),
       (2, 'Pink Dogwood', 5, 575),
       (3, 'Flowering Peach', 4, 350),
       (4, 'Flowering Crab Apple', 3, 240),
       (5, 'Magnolia', 4, 495),
       (6, 'Juniper', 3, 180),
       (7, 'Pampas', 4, 450),
       (8, 'Yellow Jasmine', 5, 570),
       (9, 'Carolina Cherry', 4, 460),
       (10, 'Camellia', 4, 495),
       (11, 'White Dogwood', 4, 520),
       (12, 'Golden Bell', 3, 155),
       (13, 'Azalea', 5, 510),
       (14, 'Chinese Fir', 4, 440),
       (15, 'Firethorn', 5, 550),
       (16, 'Redbud', 3, 170),
       (17, 'Nandina', 4, 440),
       (18, 'Holly', 4, 465);


create table masters_score_holes
(
    player_id   int references golfers (player_id),
    hole_number int references masters_holes (hole_number),
    score       int         not null,
    round       int         not null,
    ts          timestamptz not null default timezone('utc', now())
);
create index on masters_score_holes (player_id);
create index on masters_score_holes (player_id, round);
SELECT create_hypertable('masters_score_holes', 'ts', chunk_time_interval => INTERVAL '5 minutes');

create table masters_scores
(
    player_id int references golfers (player_id),
    score     int         not null,
    round     int         not null,
    ts        timestamptz not null default timezone('utc', now())
);

create index on masters_scores (player_id);
create index on masters_scores (player_id, round);
SELECT create_hypertable('masters_scores', 'ts', chunk_time_interval => INTERVAL '5 minutes');


create table rest_telemetry
(
    ip                 text,
    http_method        citext,
    url_path           citext,
    ts                 timestamptz default timezone('utc', now()),
    http_code          int,
    http_written       int,
    http_duration      double precision,
    http_duration_text text,
    user_agent_family  citext,
    user_agent_major   citext,
    user_agent_minor   citext,
    user_agent_patch   citext,
    os_family          citext,
    os_major           citext,
    os_minor           citext,
    os_patch           citext,
    os_patch_minor     citext,
    device_family      citext,
    device_brand       citext,
    device_model       citext
);

select create_hypertable('rest_telemetry', 'ts', chunk_time_interval => INTERVAL '30 minutes');

