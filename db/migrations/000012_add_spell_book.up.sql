CREATE TABLE IF NOT EXISTS spell_book
(
    id          int NOT NULL AUTO_INCREMENT,
    campaign_id int          not null,
    character_id int          not null,
    spell_stats varchar(5000) DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (campaign_id) REFERENCES campaign(id) ON DELETE CASCADE,
    FOREIGN KEY (character_id) REFERENCES `character`(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS spell_book_entry
(
    id            int NOT NULL AUTO_INCREMENT,
    spell_book_id int NOT NULL,
    spell_id      int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (spell_book_id) REFERENCES spell_book (id) ON DELETE CASCADE,
    FOREIGN KEY (spell_id) REFERENCES spell (id) ON DELETE CASCADE
);
