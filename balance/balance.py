def balance(start_balance, top_up, interest_rate):
    current_balance = start_balance + top_up
    return current_balance + start_balance * interest_rate/100

def calculate_balance(start_balance, top_up, interest_rate, n_years):
    # calculate balance for 1st year
    calculated_balance = balance(start_balance, top_up, interest_rate)

    for _ in range(2, n_years + 1):
        calculated_balance = balance(calculated_balance, top_up, interest_rate)

    return calculated_balance

if __name__ == '__main__':
    print(calculate_balance(1000, 200, 2, 3))
