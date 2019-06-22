ALTER TABLE entry ADD len MEDIUMINT;
ALTER TABLE entry ADD INDEX len_idx_on_entry(len);

UPDATE entry e1, entry e2 SET e1.len = character_length(e2.keyword) where e1.id = e2.id;
