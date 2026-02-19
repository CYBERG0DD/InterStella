# A simple, command-line to-do list application.

def display_menu():
    """Displays the main menu options to the user."""
    print("\n--- Simple To-Do List ---")
    print("1. View all tasks")
    print("2. Add a new task")
    print("3. Mark a task as completed")
    print("4. Remove a task")
    print("5. Exit")
    print("***********************\n")


def view_tasks(tasks):
    """
    Displays all tasks in the list.

    Args:
        tasks (list): A list of tuples, where each tuple contains a task string
                      and a boolean for its completion status.
    """
    print("\n--- Your Tasks ---")
    if not tasks:
        print("You have no tasks in your list.")
    else:
        for i, (task, completed) in enumerate(tasks):
            # Checkbox display: [x] for completed, [ ] for not completed
            status = "x" if completed else " "
            print(f"{i + 1}. [{status}] {task}")
    print("******************\n")


def add_task(tasks):
    """
    Prompts the user to add a new task and adds it to the list.

    Args:
        tasks (list): The list of tasks to add the new task to.
    """
    task = input("Enter the new task: ")
    if task:
        # Add as a tuple: (task_text, completed_status)
        tasks.append((task, False))
        print(f"Task '{task}' added successfully.")
    else:
        print("Task cannot be empty. Please try again.")


def mark_completed(tasks):
    """
    Prompts the user to mark a task as completed.

    Args:
        tasks (list): The list of tasks to modify.
    """
    view_tasks(tasks)
    try:
        task_num = int(
            input("Enter the number of the task to mark as completed: "))
        if 1 <= task_num <= len(tasks):
            task_index = task_num - 1
            task, _ = tasks[task_index]
            # Update the tuple with True for completed
            tasks[task_index] = (task, True)
            print(f"Task '{task}' marked as completed.")
        else:
            print("Invalid task number. Please try again.")
    except ValueError:
        print("Invalid input. Please enter a number. ")


def remove_task(tasks):
    """
    Prompts the user to remove a task from the list.

    Args:
        tasks (list): The list of tasks to modify.
    """
    view_tasks(tasks)
    try:
        task_num = int(input("Enter the number of the task to remove: "))
        if 1 <= task_num <= len(tasks):
            removed_task = tasks.pop(task_num - 1)
            print(f"Task '{removed_task[0]}' removed successfully.")
        else:
            print("Invalid task number. Please try again.")
    except ValueError:
        print("Invalid input. Please enter a number.")


def main():
    """Main function to run the to-do list application."""
    tasks = []  # The main list to store all tasks
    while True:
        display_menu()
        choice = input("Enter your choice (1-5): ")

        if choice == '1':
            view_tasks(tasks)
        elif choice == '2':
            add_task(tasks)
        elif choice == '3':
            mark_completed(tasks)
        elif choice == '4':
            remove_task(tasks)
        elif choice == '5':
            print("Exiting application. Goodbye!!!. ")
            break
        else:
            print("Invalid. Please enter a number. ")


if __name__ == "__main__":
    main()
