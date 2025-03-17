#include <iostream>
#include <fstream>
#include <vector>
#include <string>

using namespace std;

class Student
{
public:
    int id;
    string name;
    vector<string> courses;

    Student(int studentId, string studentName) : id(studentId), name(studentName) {}

    void enroll(string course)
    {
        courses.push_back(course);
    }

    void display() const
    {
        cout << "ID: " << id << " | Name: " << name << " | Courses: ";
        for (const auto &course : courses)
        {
            cout << course << " ";
        }
        cout << endl;
    }
};

class Database
{
private:
    vector<Student> students;

public:
    void addStudent(int id, string name)
    {
        students.push_back(Student(id, name));
    }

    void enrollStudent(int id, string course)
    {
        for (auto &student : students)
        {
            if (student.id == id)
            {
                student.enroll(course);
                return;
            }
        }
        cout << "Student not found!" << endl;
    }

    void displayStudents()
    {
        for (const auto &student : students)
        {
            student.display();
        }
    }

    void saveToFile(string filename)
    {
        ofstream file(filename);
        for (const auto &student : students)
        {
            file << student.id << " " << student.name << " ";
            for (const auto &course : student.courses)
            {
                file << course << " ";
            }
            file << "\n";
        }
        file.close();
    }

    void loadFromFile(string filename)
    {
        ifstream file(filename);
        if (!file)
        {
            cout << "No saved data found!" << endl;
            return;
        }
        students.clear();
        int id;
        string name, course;
        while (file >> id >> name)
        {
            Student student(id, name);
            while (file.peek() != '\n' && file >> course)
            {
                student.enroll(course);
            }
            students.push_back(student);
        }
        file.close();
    }
};

int main()
{
    Database db;
    int choice;

    do
    {
        cout << "\nStudent Management System\n";
        cout << "1. Add Student\n2. Enroll Student in Course\n3. Display Students\n4. Save & Exit\nEnter choice: ";
        cin >> choice;

        if (choice == 1)
        {
            int id;
            string name;
            cout << "Enter ID: ";
            cin >> id;
            cout << "Enter Name: ";
            cin >> name;
            db.addStudent(id, name);
        }
        else if (choice == 2)
        {
            int id;
            string course;
            cout << "Enter Student ID: ";
            cin >> id;
            cout << "Enter Course Name: ";
            cin >> course;
            db.enrollStudent(id, course);
        }
        else if (choice == 3)
        {
            db.displayStudents();
        }
    } while (choice != 4);

    db.saveToFile("students.txt");
    cout << "Data saved successfully!" << endl;
    return 0;
}
