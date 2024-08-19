-- Insert test users
INSERT INTO users (UserId) VALUES (DEFAULT), (DEFAULT), (DEFAULT);

-- Insert test logbooks
INSERT INTO logbooks (Title) VALUES
('Personal Log 1'),
('Personal Log 2'),
('Personal Log 3');

-- Insert test entries
INSERT INTO entries (Title, Description, CreatedBy) VALUES
('Baked a pie', 'Bought ingredients at store, combined them, and baked them into a nice pie.', 1),
('Completed project report', 'Compiled data, wrote the report, and submitted it to the manager.', 1),
('Went for a run', 'Ran 5 miles in the park, enjoyed the fresh air and exercise.', 1),
('Read a book', 'Finished reading "The Great Gatsby". Took notes on key themes and characters.', 2),
('Cleaned the house', 'Vacuumed, dusted, and organized various rooms in the house.', 2),
('Grocery shopping', 'Purchased groceries for the week including fruits, vegetables, and dairy products.', 2),
('Attended team meeting', 'Discussed project milestones, set new goals, and assigned tasks.', 3),
('Cooked dinner', 'Prepared a homemade lasagna with a side salad and garlic bread.', 3),
('Watched a movie', 'Saw "Inception" at the theater, enjoyed the plot and special effects.', 3),
('Completed online course', 'Finished an online course on data science and received a completion certificate.', 3);

-- Associate entries with logbooks
INSERT INTO belongs_to (EntryId, LogbookId) VALUES
(1, 1), -- Baked a pie
(2, 1), -- Completed project report
(3, 1), -- Went for a run
(4, 2), -- Read a book
(5, 2), -- Cleaned the house
(6, 2), -- Grocery shopping
(7, 3), -- Attended team meeting
(8, 3), -- Cooked dinner
(9, 3), -- Watched a movie
(10, 3); -- Completed online course

-- Associate logbooks with users
INSERT INTO owned_by (LogbookId, UserId) VALUES
(1, 1),
(2, 2),
(3, 3);
