import React, { FC, MouseEventHandler } from 'react';
import styles from './styles.module.css';

interface Props {
  height?: number;
  width?: number;
  className?: string;
  onClick?: MouseEventHandler<HTMLImageElement>;
}

enum IconType {
  CHECKED = 'check_box',
  BLANK = 'check_box_outline_blank'
}

export const CheckedIcon: FC<Props> = ({
  height = 20,
  width = 20,
  className,
  onClick
}) => {
  return (
    <img
      src={`/icons/${IconType.CHECKED}.svg`}
      alt='check box'
      height={height}
      width={width}
      onClick={onClick}
      className={[styles.clickable, className].join(' ')}
    />
  )
}

export const BlankIcon: FC<Props> = ({
  height = 20,
  width = 20,
  className,
  onClick
}) => {
  return (
    <img
      src={`/icons/${IconType.BLANK}.svg`}
      alt='check box'
      height={height}
      width={width}
      onClick={onClick}
      className={[styles.clickable, className].join(' ')}
    />
  )
}

export default CheckedIcon;