#include <iostream>
#include <string>

class Student{
    public:
    int age();
    void showAge();
};

int Student::age(){
    return 10;
}

void Student::showAge(){
    std::cout << "age: 10" << std::endl;
}

void _showAge(){
    Student().showAge();
}
