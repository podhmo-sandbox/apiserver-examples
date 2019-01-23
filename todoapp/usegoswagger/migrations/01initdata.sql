insert into users (id, name) values (1, 'foo') on conflict (id) do nothing;
