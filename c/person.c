// person.c
#include "person.h"
#include <string.h>

int person_is_adult(struct person p)
{
    // Logic to determine if the person is an adult
    return (p.age >= 18) ? 1 : 0;
}

int person_is_male(struct person p)
{
    // Compare the sex field with "male"
    return (strcmp(p.sex, "male") == 0) ? 1 : 0;
}