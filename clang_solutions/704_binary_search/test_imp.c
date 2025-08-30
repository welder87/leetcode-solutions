#include "imp.h"
#include <assert.h>
#include <stdio.h>

typedef struct {
  int *nums;
  size_t cnt;
  int target;
  int expectedValue;
} TestCase;

void print_test_case(TestCase *tc) {
  printf("numsSize: %zu\n", tc->cnt);
  printf("target: %d\n", tc->target);
  printf("expectedValue: %d\n", tc->expectedValue);
  printf("nums: ");
  for (size_t i = 0; i < tc->cnt; i++) {
    printf("%d ", tc->nums[i]);
  }
  printf("\n");
}

void run_test_cases(TestCase *test_cases, size_t count,
                    int func(int *nums, int numsSize, int target)) {
  size_t j = 1;
  for (size_t i = 0; i < count; i++) {
    printf("===============================================================\n");
    printf("Test Case: %zu\n", i + j);
    printf("===============================================================\n");
    print_test_case(&test_cases[i]);
    int cnt = (int)test_cases[i].cnt;
    assert(func(test_cases[i].nums, cnt, test_cases[i].target) ==
           test_cases[i].expectedValue);
  }
}

int main() {
  int array0[] = {-1, 0, 3, 5, 9, 12};
  int array1[] = {-36, -6, -1, 0, 1, 5, 9, 12};
  int array2[] = {-1};
  int array3[] = {0};
  size_t num_sizes[] = {6, 8, 1, 1};

  size_t test_case_size = 11;
  TestCase test_cases[] = {
      // preset cases
      {
          .nums = array0,
          .cnt = num_sizes[0],
          .target = 9,
          .expectedValue = 4,
      },
      {
          .nums = array0,
          .cnt = num_sizes[0],
          .target = 2,
          .expectedValue = -1,
      },
      // common cases
      {
          .nums = array1,
          .target = 2,
          .cnt = num_sizes[1],
          .expectedValue = -1,
      },
      {
          .nums = array1,
          .target = -7,
          .cnt = num_sizes[1],
          .expectedValue = -1,
      },
      {
          .nums = array1,
          .target = -36,
          .cnt = num_sizes[1],
          .expectedValue = 0,
      },
      {
          .nums = array1,
          .target = -6,
          .cnt = num_sizes[1],
          .expectedValue = 1,
      },
      {
          .nums = array1,
          .target = 0,
          .cnt = num_sizes[1],
          .expectedValue = 3,
      },
      // corner cases
      {
          .nums = array2,
          .target = -1,
          .cnt = num_sizes[2],
          .expectedValue = 0,
      },
      {
          .nums = array2,
          .target = 0,
          .cnt = num_sizes[2],
          .expectedValue = -1,
      },
      {
          .nums = array2,
          .target = 1,
          .cnt = num_sizes[2],
          .expectedValue = -1,
      },
      {
          .nums = array3,
          .target = 0,
          .cnt = num_sizes[3],
          .expectedValue = 0,
      },
  };
  run_test_cases(test_cases, test_case_size, search);
  return 0;
}
