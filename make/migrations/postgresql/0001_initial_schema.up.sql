create table access (
 access_id SERIAL PRIMARY KEY NOT NULL,
 access_code char(1),
 comment varchar (30)
);

insert into access (access_code, comment) values 
('M', 'Management access for project'),
('R', 'Read access for project'),
('W', 'Write access for project'),
('D', 'Delete access for project'),
('S', 'Search access for project');