INSERT INTO author(name) VALUES
                             ('Brian Kernighan'),
                             ('Dennis Ritchie'),
                             ('Alan A. A. Donovan'),
                             ('Stephen King'),
                             ('Harper Lee'),
                             ('Jane Austen'),
                             ('J.K. Rowling'),
                             ('Aleko Konstantinov');

INSERT INTO book(title) VALUES
                            ('Misery'),
                            ('The Shining'),
                            ('To Kill a Mockingbird'),
                            ('Harry Potter and the Order of the Phoenix'),
                            ('Chicago and Back'),
                            ('The C Programming Language'),
                            ('The Go Programming Language'),
                            ('Pride and Prejudice');

INSERT INTO book_author(author_id, book_id) VALUES
                                                ('1','6'),
                                                ('1','7'),
                                                ('2','6'),
                                                ('3','7'),
                                                ('4','1'),
                                                ('4','2'),
                                                ('5','3'),
                                                ('6','8'),
                                                ('7','4'),
                                                ('8','5');