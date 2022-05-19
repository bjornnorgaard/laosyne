CREATE TABLE pictures
(
    id        UUID PRIMARY KEY,
    path      TEXT      NOT NULL,
    ext       TEXT      NOT NULL,
    views     INT       NOT NULL,
    likes     INT       NOT NULL,
    rating    FLOAT     NOT NULL,
    deviation FLOAT     NOT NULL,
    wins      INT       NOT NULL,
    losses    INT       NOT NULL,
    created   TIMESTAMP NOT NULL,
    updated   TIMESTAMP NOT NULL
);

CREATE TABLE media_paths
(
    id      UUID PRIMARY KEY,
    path    TEXT      NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL
);
