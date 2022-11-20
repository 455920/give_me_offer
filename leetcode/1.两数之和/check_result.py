from solution import Solution


def check_result(l1, l2):
    list.sort(l1)
    list.sort(l2)
    if len(l1) != len(l2):
        return False
    for i in range(0, len(l1)):
        if l1[i] != l2[i]:
            return False
    return True


if __name__ == "__main__":
    with open('test_data.txt', 'r') as fd:
        for line in fd.readlines():
            line = line.strip()
            test_data_vec = line.split('|')
            if len(test_data_vec) != 3:
                continue
            input_vec = test_data_vec[0].split(',')
            input_vec = [int(i) for i in input_vec]
            input_target = int(test_data_vec[1])
            expect_result = test_data_vec[2].split(',')
            expect_result = [int(i) for i in expect_result]

            print(
                f"input_vec: {input_vec}, input_target: {input_target}, expect_result: {expect_result}")
            actual_result = Solution().twoSum(input_vec, input_target)
            res = check_result(actual_result, expect_result)
            print(f"actual_result: {actual_result}, check_result: {res}")
            print(
                "------------------------------------------------------------------------")
