CREATE TABLE IF NOT EXISTS events (
    `EventType` String,
    `Time` DateTime DEFAULT NOW(),
    `UserIP` String
)
ENGINE = MergeTree
PARTITION BY toYYYYMM(Time)
ORDER BY (EventType, UserIP, Time);
