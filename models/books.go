package models

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func GetAllBooks() ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}

func GetBook(isbn string) (*Book, error) {
  row := db.QueryRow("SELECT * FROM books WHERE isbn = ?", isbn)
  bk := new(Book)
  err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
    return bk,err
}

func CreateBook(isbn string, title string, author string, price float64) (int64, error) {
  result, err := db.Exec("INSERT INTO books VALUES(?, ?, ?, ?)", isbn, title, author, price)
  if err != nil {
        return 0, err
    }
  rowsAffected, err := result.RowsAffected()
   if err != nil {
        return 0, err
    }
   return rowsAffected,err
}

func DeleteBook(isbn string) (int64, error) {
  result, err := db.Exec("DELETE FROM books WHERE isbn = ?", isbn)
  if err != nil {
        return 0, err
    }
  rowsAffected, err := result.RowsAffected()
   if err != nil {
        return 0, err
    }
   return rowsAffected,err
}
