import React , { FC, MouseEventHandler, useState } from 'react';
import styles from './styles.module.css';
import Card from '../../atoms/Card';
import CheckBoxIcon from '../../molecules/CheckBox';
import { Todo }  from '../../common/Type'

interface PresenterProps {
  title: string;
  isDone: boolean;
  handleToggleIcon?: MouseEventHandler<HTMLImageElement>;
  className?: string;
}

const TodoItemPresenter: FC<PresenterProps> = ({
  title,
  isDone,
  handleToggleIcon
}) => {
  return (
    <Card className={styles.root}>
      <CheckBoxIcon
        state={isDone}
        onClick={handleToggleIcon}
        className={styles.icon}
      />
      <p className={styles.title}>{title}</p>
    </Card>
  )
}

interface ContainerProps {
  presenter: typeof TodoItemPresenter;
  todo: Todo;
}

const TodoItemContainer: FC<ContainerProps> = ({
  presenter,
  todo,
}) => {
  const [isDone, setState] = useState(todo.done);
  const handleToggleIcon = () => {
    setState(!isDone);
    console.log(isDone);
  }
  const title = todo.title;
  const presenterProps = { title, isDone, handleToggleIcon }
  return (
    presenter(presenterProps)
  )
}

interface Props {
  todo: Todo;
}

const TodoItem: FC<Props> = ({ todo }) => {
  return (
    <TodoItemContainer
      presenter={presenterProps => <TodoItemPresenter {...presenterProps}/>}
      todo={todo}
    />
  )
}

export default TodoItem;