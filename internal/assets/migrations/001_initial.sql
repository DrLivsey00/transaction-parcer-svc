-- +migrate Up

create table transfers(
    id serial primary key,
    tx_hash varchar(66) not null,
    sender varchar(42) not null,
    receiver varchar(42) not null,
    token_amount real not null
);

create index receiver_index on transfers(receiver);
create index sender_index on transfers(sender);

-- +migrate Down
drop table transfers;