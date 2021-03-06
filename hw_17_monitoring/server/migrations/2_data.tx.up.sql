insert into calendar.calendar (date, description)
values ('2019-11-10', 'some desc'),
       ('2019-11-12', 'some desc2'),
       ('2019-11-15', 'some desc3'),
       ('2019-10-20', 'some desc4')
       ;

insert into calendar.event (date_fk, time, title, description)
values
       ((SELECT id FROM calendar.calendar WHERE date = '2019-11-10'), '2019-11-10 07:18:09.767953 +00:00', 'some title event1', 'desc event1'),
       ((SELECT id FROM calendar.calendar WHERE date = '2019-11-10'), '2019-11-10 09:20:09.767953 +00:00', 'some title event2', 'desc event2'),
       ((SELECT id FROM calendar.calendar WHERE date = '2019-11-12'), '2019-11-12 10:20:09.767953 +00:00', 'some title event3', 'desc event3'),
       ((SELECT id FROM calendar.calendar WHERE date = '2019-10-20'), '2019-10-20 10:30:09.767953 +00:00', 'some title event4', 'desc event4')
       ;
