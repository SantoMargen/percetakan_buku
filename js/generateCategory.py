# Python Script to Generate .sql File

with open("insert_categories.sql", "w") as f:
    f.write("-- Automatically generated insert statements\n")
    f.write("INSERT INTO \"public\".\"category\" (\"category_id\", \"category_name\", \"description\", \"entry_user\", \"entry_time\") VALUES\n")
    
    for i in range(1, 10001):
        if i < 10000:
            f.write(f"({i}, 'Ilussion', '123 Main St, City A, Country A', 'Hisbikal', '2024-11-01 11:34:27.165961'),\n")
        else:
            f.write(f"({i}, 'Ilussion', '123 Main St, City A, Country A', 'Hisbikal', '2024-11-01 11:34:27.165961');\n")

print("SQL file generated successfully!")
