import React, { FC, MouseEventHandler } from 'react';
import styles from './styles.module.css';
import { CheckedIcon, BlankIcon } from '../../atoms/Icon';

interface PresenterProps {
  icon: typeof CheckedIcon | typeof BlankIcon;
  onClick?: MouseEventHandler<HTMLImageElement>;
  className?: string;
}

const CheckBoxIconPresenter: FC<PresenterProps> = ({
  icon:Tag,
  onClick,
  className
}) => {
  return (
    <Tag
      onClick={onClick}
      className={[styles.root, className].join(' ')}
    />
  )
}

interface ContainerProps {
  presenter: typeof CheckBoxIconPresenter;
  isChecked: boolean;
  onClick?: MouseEventHandler<HTMLImageElement>;
}

const CheckBoxIconContainer: FC<ContainerProps> = ({
  presenter,
  isChecked,
  onClick
}) => {
  const icon = isChecked ?
    (props: any) => <CheckedIcon {...props}/>:
    (props: any) => <BlankIcon {...props}/>
  const presenterProps = { icon, onClick }
  return (
    presenter(presenterProps)
  )
}

interface Props {
  onClick?: MouseEventHandler<HTMLImageElement>;
  state: boolean;
}

const CheckBoxIcon: FC<Props> = ({ onClick, state }) => {
  return (
    <CheckBoxIconContainer
      presenter={presenterProps => <CheckBoxIconPresenter {...presenterProps}
      />}
      isChecked={state}
      onClick={onClick}
    />
  )
}

export default CheckBoxIcon;