// person.h
struct person
{
    int  age;
    char *sex;
    // Other fields of the person struct
};

int person_is_adult(struct person p);
int person_is_male(struct person p);