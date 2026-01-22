ALTER TABLE `character` DROP CONSTRAINT character_ibfk_2;
ALTER TABLE `character` DROP COLUMN spell_book_id;

TRUNCATE TABLE spell_book_entry;
DROP TABLE spell_book_entry;

TRUNCATE TABLE spell_book;
DROP TABLE spell_book;
