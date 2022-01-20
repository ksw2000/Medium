#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct pair {
    char* key;
    int val;
    struct pair* next;
} Pair;

typedef struct hashmap {
    Pair** list;       // Pair* list
    unsigned int cap;  // capacity, the length of list
    unsigned int len;  // length, the number of pairs
} HashMap;

HashMap* newHashMap() {
    HashMap* this = malloc(sizeof(this));
    this->cap = 8;  // set default capacity
    this->len = 0;  // no pair in map
    // set all pointer to null in this->list
    this->list = calloc((this->cap), sizeof(Pair*));
    return this;
}

unsigned hashcode(HashMap* this, char* key) {
    unsigned code;
    for (code = 0; *key != '\0'; key++) {
        code = *key + 31 * code;
    }
    return code % (this->cap);
}

int get(HashMap* this, char* key) {
    Pair* current;
    for (current = this->list[hashcode(this, key)]; current;
         current = current->next) {
        if (!strcmp(current->key, key)) {
            return current->val;
        }
    }
    fprintf(stderr, "%s is not found\n", key);
    exit(EXIT_FAILURE);
}

// If key is not in hashmap, put into map. Otherwise, replace it.
void set(HashMap* this, char* key, int val) {
    unsigned index = hashcode(this, key);
    Pair* current;
    for (current = this->list[index]; current; current = current->next) {
        // if key has been already in hashmap
        if (!strcmp(current->key, key)) {
            current->val = val;
            return;
        }
    }

    // key is not in hashmap
    Pair* p = malloc(sizeof(*p));
    p->key = key;
    p->val = val;
    p->next = this->list[index];
    this->list[index] = p;
    this->len++;
}

int main() {
    HashMap* m = newHashMap();
    // set value
    set(m, "Ayumu", 1);
    set(m, "Honoka", 2);
    set(m, "Chika", 3);
    set(m, "Chisato", 4);

    // get value
    printf("Ayumu->%d\n", get(m, "Ayumu"));
    printf("Honoka->%d\n", get(m, "Honoka"));
    printf("Chika->%d\n", get(m, "Chika"));
    printf("Chisato->%d\n", get(m, "Chisato"));

    // replace value
    set(m, "Ayumu", 10);
    printf("Ayumu->%d\n", get(m, "Ayumu"));
    printf("Map cap:%u len:%u\n", m->cap, m->len);
    return 0;
}