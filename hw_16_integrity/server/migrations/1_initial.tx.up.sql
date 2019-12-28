CREATE OR REPLACE FUNCTION upd_updated_at() RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;

create or replace function add_time_fields(table_name text) returns void
    language plpgsql
as
$$
declare
    trigger_name text;
begin
    execute 'alter table ' || table_name || ' add column created_at timestamp with time zone default now() not null;';
    execute 'alter table ' || table_name || ' add column updated_at timestamp with time zone default now() not null;';
    trigger_name := 't_' || table_name || '_upt';

    execute 'create trigger ' || trigger_name || ' before update on ' || table_name ||
            ' for each row execute procedure upd_updated_at()';
end;
$$;

CREATE TABLE calendar
(
    id             SERIAL PRIMARY KEY,
    date           DATE    NOT NULL,
    description    TEXT,
    is_celebration BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT public_date_date_uidx UNIQUE (date)
);
SELECT add_time_fields('calendar');

CREATE TABLE event
(
    id          BIGSERIAL PRIMARY KEY,
    date_fk     INT  NOT NULL,
    time        TIMESTAMP WITH TIME ZONE,
    title       TEXT NOT NULL,
    description TEXT,
    CONSTRAINT public_event_time_date_uidx UNIQUE (time, date_fk),
    CONSTRAINT public_event_date_calendar_fk FOREIGN KEY (date_fk) REFERENCES calendar (id)
);
SELECT add_time_fields('event');

