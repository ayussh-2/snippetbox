CREATE DATABASE snippetbox

USE snippetbox

CREATE Table snippets (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
)


INSERT INTO snippets (title, content, created, expires)
VALUES (
    'An old silent pond',
    'An old silent pond ...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Basho',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);


INSERT INTO snippets (title, content, created, expires)
VALUES (
    'Over the wintry forest',
    'Over the wintry forest,
winds howl in rage
with no leaves to blow.

- Natsume Soseki',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);


INSERT INTO snippets (title, content, created, expires)
VALUES (
    'First autumn morning',
    'First autumn morning
the mirror I stare into
shows my father''s face.

- Murakami Kijo',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

CREATE USER 'web'@'%';

GRANT SELECT, INSERT, UPDATE, DELETE, CREATE ON snippetbox.* TO 'web'@'%';

ALTER USER 'web'@'%' IDENTIFIED BY 'pass';

SELECT * FROM snippets;