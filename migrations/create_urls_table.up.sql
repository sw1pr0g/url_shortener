CREATE TABLE urls
(
    uuid      UUID DEFAULT gen_random_uuid() NOT NULL,
    url       TEXT                           NOT NULL
        CONSTRAINT unique_url
            UNIQUE,
    short_url text                           NOT NULL
        CONSTRAINT unique_short_url
            UNIQUE
);