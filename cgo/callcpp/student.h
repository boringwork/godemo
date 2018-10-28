#ifndef __STUDENT_H__
#define __STUDENT_H__

template <typename T> static constexpr bool SkToBool(const T& x) { return 0 != x; }

class Student{
public:
    void display();
};


#endif//__STUDENT_H__