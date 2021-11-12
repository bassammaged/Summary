'''
    # ----- start connection
        sqlite3.connect()
        - This API opens a connection to the SQLite database file.

    # ----- execute query
        cursor.execute() 
        
    # ----- Commit the queries
        obj1.commit()

    # ----- CLose the connection
        obj1.close()

    V.I Note: 
        You need a cursor object to fetch results. 
'''
import sqlite3

# ----- Craete a table
def createConnDB():
    try:
        conn    = sqlite3.connect('Secranix-Co.db')
        print("Opened database successfully.")

        return conn 
    except:
        print("Something goes wrong while open the db.")


# ----- Create table
def CreateTable(conn):
    try:
        cur     = conn.cursor()
        query   = '''CREATE TABLE COMPANY
            (ID INT PRIMARY KEY     NOT NULL,
            NAME           TEXT    NOT NULL,
            AGE            INT     NOT NULL,
            ADDRESS        CHAR(50),
            SALARY         REAL);'''

        cur.execute(query)
        print("The table has been created successfully.")
    except:
        print("Something goes wrong while creating the table.")
    finally:
        cur.close()
    

# ----- Show the table    
def showTable(conn):
    query   = "SELECT name FROM sqlite_master WHERE type='table';"
    cur     = conn.cursor()
    try:
        cur.execute(query)
        print("Table has been printed.")
        return cur.fetchall() 
    except:
        print("Something goes wrong while showing the content of table.")
    finally:
        cur.close()


# ----- Insert Values:
def insertValues(conn):
    try:
        conn.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESS,SALARY) \
        VALUES (1, 'Paul', 32, 'California', 20000.00 )");

        conn.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESS,SALARY) \
            VALUES (2, 'Allen', 25, 'Texas', 15000.00 )");

        conn.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESS,SALARY) \
            VALUES (3, 'Teddy', 23, 'Norway', 20000.00 )");

        conn.execute("INSERT INTO COMPANY (ID,NAME,AGE,ADDRESS,SALARY) \
            VALUES (4, 'Mark', 25, 'Rich-Mond ', 65000.00 )");

        conn.commit()
    except:
        print("Something goes wrong while inserting new content of table")
    finally:
        conn.close()

def main():
    try:
        conn    = createConnDB()
        CreateTable(conn)

        conn    = createConnDB()
        result  = showTable(conn)
        print(result)
    
        conn    = createConnDB()
        insertValues(conn)
    except:
        print("Somethind goes wrong!")



if __name__ == "__main__":
    main()