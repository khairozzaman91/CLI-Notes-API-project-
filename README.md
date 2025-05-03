# CLI Notes API

This is a simple **CLI application** where users can manage notes through the terminal. It allows users to perform the following operations:

- **Create Notes**
- **Update Notes**
- **Show All Notes**
- **Delete Notes**

The application uses **SQLite** for local storage, and **GORM** as the ORM for database interaction.

---

## Features

### CRUD Operations:

- **Create**: Add new notes with a title and content.
- **Read**: View a list of all saved notes.
- **Update**: Modify the title and content of an existing note.
- **Delete**: Remove a note from the database.

### User-Friendly Command-Line Interface:
Easy-to-follow menu and prompts.

### SQLite Database:
Stores notes locally for persistence.

### Go & GORM:
Leveraging Go for backend logic and GORM for interacting with the database.

---

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/cli-notes-api.git
   cd cli-notes-api
