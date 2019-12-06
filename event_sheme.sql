create table events (
 id serial primary key,
 owner bigint,
 title text,
 description text,
 start_time time,
 end_time time
);
create index owner_idx on events (owner);
create index start_idx on events using btree (start_time);