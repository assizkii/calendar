create table if not exists events (
    id serial primary key,
    owner_id int,
    title text,
    description text,
    start timestamp not null,
    end_time timestamp
);
create index if not exists owner_idx on events (owner_id);
create index if not exists start_idx on events using btree (start, end_time);