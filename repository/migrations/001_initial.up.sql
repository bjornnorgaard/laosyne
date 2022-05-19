CREATE TABLE IF NOT EXISTS pictures
(
    id        BIGSERIAL PRIMARY KEY,
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

CREATE TABLE IF NOT EXISTS paths
(
    id      BIGSERIAL PRIMARY KEY,
    path    TEXT      NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL
);
