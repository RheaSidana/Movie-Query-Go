<h1>Movie Query Go Application</h1>

<h3>Steps to run the application</h3>
<br />

<h2>
  1. Clone the repo
  <br /><br />
  2. run the command: 
</h2>

``` 
go mod tidy 
```

<br />
<h2>
  3. DB operations
</h2>
<h4>
  <br /><br />
  &emsp; <<<< INSTALLING >>>> <br />
  &emsp;&emsp; a. MAC
  <br />
  &emsp;&emsp;&emsp;&emsp; 1. install postgres: 
</h4>

``` 
brew install postgresql 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h4>2. start/stop postgres service: <br /></h4>

``` 
brew services start/stop postgresql 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h4>3. <br /></h4>

``` 
psql postgres 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h4>4. add postgres password: <br /></h4>

``` 
\password {password}; 
```

<br /><br />&emsp;&emsp;
<h4>
  b. Windows
  <br />
  &emsp;&emsp;&emsp;&emsp; 1. install postgres:
  <!-- (https://www.postgresql.org/download/windows/) -->
  <a href="https://www.postgresql.org/download/windows/">[Link text Here]</a>
  <br />
  &emsp;&emsp;&emsp;&emsp; 2. port: 5432 (defaut), user: postgres (defaut)
  <br />
  &emsp;&emsp;&emsp;&emsp; 3. add postgres password
  <br />
  &emsp;&emsp;&emsp;&emsp; 4. open psql sql shell <br /><br />
  &emsp; <<<< CREATING >>>> <br />&emsp;&emsp;&emsp;&emsp; 5. create orm db: <br />
</h4>

``` 
CREATE database movies_app; 
```

<br /><br />&emsp;
<h4>
  <<<< CONNECTING >>>>
  <br />&emsp;&emsp;&emsp;&emsp; 6. connect to db: <br />
</h4>
``` 
\c "db" 
```

<br />&emsp;&emsp;&emsp;&emsp;
<h4>
  7. Edit .env file with postgres details
  <br /><br /><br />
  4. Migrate Tables <br />
</h4>

``` 
go run .\migrations\migrate.go 
```

<br /><br />&emsp;
<h4>View DB Table Schemas: <br /></h4>

``` 
\d "tablename" 
```

<br /><br />
<h4>5. Seed Data to the Table <br /></h4>

``` 
go run .\dataSeeding\seedData.go 
```

<br /><br />
<h4>6. Run the application <br /></h4>

``` 
go run . 
```

<br /><br />
<h4>
  7. Call APIs
  <!-- https://docs.google.com/document/d/1-6-EVSzGkoIPzp5y3kb0rDhq1YrrqM9qdAXzhDRCEKg/edit?usp=sharing -->
  <a
    href="https://docs.google.com/document/d/1-6-EVSzGkoIPzp5y3kb0rDhq1YrrqM9qdAXzhDRCEKg/edit?usp=sharing"
    >[API DOC]</a
  >
</h4>
