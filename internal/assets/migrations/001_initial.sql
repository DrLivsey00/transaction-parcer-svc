-- +migrate Up

create table transfers(
    tx_hash text primary key,
    sender text not null,
    receiver text not null,
    amount decimal not null,
);

-- +migrate Down
drop table transfers