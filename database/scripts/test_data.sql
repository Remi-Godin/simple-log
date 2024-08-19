-- Insert test users
INSERT INTO users (UserId) VALUES (DEFAULT), (DEFAULT), (DEFAULT);

-- Insert test logbooks
INSERT INTO logbooks (Title,OwnedBy) VALUES
('Personal Log 1',1),
('Personal Log 2',2),
('Personal Log 3',3);

-- Insert test entries
INSERT INTO entries (Title, Description, CreatedBy, LogbookId) VALUES
('Baked a pie', 'Bought ingredients at store, combined them, and baked them into a nice pie.', 1,1),
('Completed project report', 'Compiled data, wrote the report, and submitted it to the manager.', 1,1),
('Went for a run', 'Ran 5 miles in the park, enjoyed the fresh air and exercise.', 1,1),
('Read a book', 'Finished reading "The Great Gatsby". Took notes on key themes and characters.', 2,2),
('Cleaned the house', 'Vacuumed, dusted, and organized various rooms in the house.', 2,2),
('Grocery shopping', 'Purchased groceries for the week including fruits, vegetables, and dairy products.', 2,2),
('Attended team meeting', 'Discussed project milestones, set new goals, and assigned tasks.', 3,3),
('Cooked dinner', 'Prepared a homemade lasagna with a side salad and garlic bread.', 3,3),
('Watched a movie', 'Saw "Inception" at the theater, enjoyed the plot and special effects.', 3,3),
('Completed online course', 'Finished an online course on data science and received a completion certificate.', 3,3);
