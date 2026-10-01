INSERT INTO public.persons (id, name, age) VALUES (1, 'Muftin', 25);
INSERT INTO public.persons (id, name, age) VALUES (2, 'Carlos', 24);
INSERT INTO public.persons (id, name, age) VALUES (3, 'Ridho', 24);
INSERT INTO public.persons (id, name, age) VALUES (6, 'Fajar', 35);
INSERT INTO public.persons (id, name, age) VALUES (7, 'Given', 19);
INSERT INTO public.persons (id, name, age) VALUES (8, 'Maruf', 30);
INSERT INTO public.persons (id, name, age) VALUES (9, 'Rama', 27);
INSERT INTO public.persons (id, name, age) VALUES (10, 'Habib', 38);
INSERT INTO public.persons (id, name, age) VALUES (11, 'Alfan', 40);
INSERT INTO public.persons (id, name, age) VALUES (13, 'koda', 20);
INSERT INTO public.persons (id, name, age) VALUES (12, 'Nico', 26);

SELECT setval('persons_id_seq', 13, true);