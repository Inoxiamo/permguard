-- Copyright 2024 Nitro Agility S.r.l.
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.
--
-- SPDX-License-Identifier: Apache-2.0

-- +goose Up
CREATE TABLE accounts (
    account_id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    name TEXT NOT NULL UNIQUE
);

CREATE INDEX accounts_name_idx ON accounts(name);

-- Creating the `account_changestreams` table
CREATE TABLE account_changestreams (
    changestream_id INTEGER PRIMARY KEY,
	change_type TEXT NOT NULL,
	change_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    account_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    name TEXT NOT NULL
);

CREATE INDEX account_changestreams_name_idx ON account_changestreams(name);

-- Trigger to track changes in the `accounts` table after insert
-- +goose StatementBegin
CREATE TRIGGER account_changestreams_after_insert
AFTER INSERT ON accounts
FOR EACH ROW
BEGIN
    INSERT INTO account_changestreams (change_type, account_id, created_at, updated_at, name)
    	VALUES ("INSERT", NEW.account_id, NEW.created_at, NEW.updated_at, NEW.name);
END;
-- +goose StatementEnd

-- Trigger to track changes in the `accounts` table after update
-- +goose StatementBegin
CREATE TRIGGER account_changestreams_after_update
AFTER UPDATE ON accounts
FOR EACH ROW
BEGIN
    UPDATE accounts SET updated_at = CURRENT_TIMESTAMP WHERE account_id = OLD.account_id;
    INSERT INTO account_changestreams (change_type, account_id, created_at, updated_at, name)
	    VALUES ("UPDATE", NEW.account_id, NEW.created_at, CURRENT_TIMESTAMP, NEW.name);
END;
-- +goose StatementEnd

-- Trigger to track changes in the `accounts` table after delete
-- +goose StatementBegin
CREATE TRIGGER account_changestreams_after_delete
AFTER DELETE ON accounts
FOR EACH ROW
BEGIN
    INSERT INTO account_changestreams (change_type, account_id, created_at, updated_at, name)
    	VALUES ("DELETE", OLD.account_id, OLD.created_at, OLD.updated_at, OLD.name);
END;
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS account_changestreams_after_insert;
DROP TRIGGER IF EXISTS account_changestreams_after_update;
DROP TRIGGER IF EXISTS account_changestreams_after_delete;
DROP TABLE IF EXISTS account_changestreams;
DROP TABLE IF EXISTS accounts;