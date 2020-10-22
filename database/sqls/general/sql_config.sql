create user genuser identified by "mysql123Admin@";
GRANT ALL PRIVILEGES ON morty.* TO genuser@"%" IDENTIFIED BY 'mysql123Admin@' WITH GRANT OPTION;
flush privileges;