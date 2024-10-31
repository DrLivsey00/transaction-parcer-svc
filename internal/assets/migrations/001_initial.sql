-- +migrate Up

create table transfers(
    tx_hash text primary key,
    sender text not null,
    receiver text not null,
    token_amount decimal not null,
);

create index receiver_index on transfers(receiver);
create index sender_index on transfers(sender)

-- +migrate Down
drop table transfers