-- +migrate Up

create table transfers(
    id serial primary key,
    tx_hash varchar(66) not null,
    sender varchar(42) not null,
    receiver varchar(42) not null,
    token_amount text not null,
    block_number integer not null, --enterpretation of block_height
    event_index integer not null
);

create unique index transfer_index on transfers(tx_hash,event_index) --composited unique index to define transfer`s uniquness
create index receiver_index on transfers(receiver);
create index sender_index on transfers(sender);

-- +migrate Down
drop table transfers;