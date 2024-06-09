INSERT INTO parties (id, name, slogan, opened_date, description) VALUES
('11111111-1111-1111-1111-111111111111', 'Freedom Party', 'Liberty for All', '1992-04-01', 'Dedicated to personal freedoms and rights.'),
('22222222-2222-2222-2222-222222222222', 'Green Alliance', 'Sustainable Future', '1989-11-09', 'Focusing on environmental and sustainable policies.'),
('33333333-3333-3333-3333-333333333333', 'Technocrat Front', 'Innovation and Progress', '1999-07-15', 'Promoting technology-driven policies.');

INSERT INTO publics (id, first_name, last_name, birthday, gender, nation, party_id) VALUES
('77777777-7777-7777-7777-777777777777', 'John', 'Doe', '1980-12-10', 'Male', 'Nomad', '11111111-1111-1111-1111-111111111111'),
('88888888-8888-8888-8888-888888888888', 'Alice', 'Johnson', '1992-05-16', 'Female', 'Settler', '22222222-2222-2222-2222-222222222222'),
('99999999-9999-9999-9999-999999999999', 'Michael', 'Smith', '1978-08-25', 'Male', 'Nomad', '33333333-3333-3333-3333-333333333333');