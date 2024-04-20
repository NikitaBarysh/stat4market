ATTACH TABLE _ UUID '80e06e1f-bdf0-4426-a4ba-1ec4c8071a6f'
(
    `eventID` Int64,
    `eventType` String,
    `userID` Int64,
    `eventTime` DateTime,
    `payload` String
)
ENGINE = MergeTree
ORDER BY (eventID, eventTime)
SETTINGS index_granularity = 8192
