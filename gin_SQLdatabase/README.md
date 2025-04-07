
-- to login to the mysql server from terminal
prasuna@prasuna:~$ mysql -u myuser -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 14
Server version: 8.0.41-0ubuntu0.22.04.1 (Ubuntu)

Copyright (c) 2000, 2025, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.


mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| prasuna            |
| sys                |
+--------------------+
5 rows in set (0.02 sec)


mysql> use prasuna
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed

mysql> SHOW TABLES;
+-------------------+
| Tables_in_prasuna |
+-------------------+
| users             |
+-------------------+
1 row in set (0.00 sec)


mysql> select * from users;
+----+---------+---------------------+------+
| id | name    | email               | age  |
+----+---------+---------------------+------+
|  1 | Alice   | alice@example.com   |   25 |
|  2 | Bob     | bob@example.com     |   30 |
|  3 | Charlie | charlie@example.com |   22 |
|  4 | David   | david@example.com   |   28 |
| 21 | Eve     | eve@example.com     |   26 |
+----+---------+---------------------+------+
5 rows in set (0.00 sec)

mysql> exit
Bye
prasuna@prasuna:~$