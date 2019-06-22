ALTER TABLE entry ADD len MIDIUMINT;
ALTER TABLE entry ADD INDEX len_idx_on_entry(len);

UPDATE entry e1, entry e2 SET e1.len = character_length(e2.description) where e1.id = e2.id;
