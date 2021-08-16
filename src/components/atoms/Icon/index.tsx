import React, { FC } from 'react';
import './styles.css';

interface Props {
  height?: number;
  width?: number;
  className?: string;
}

enum IconType {
  CHECKED = 'check_box',
  BLANK = 'check_box_outline_blank'
}

export const CheckedIcon: FC<Props> = ({
  height = 20,
  width = 20,
  className
}) => {
  return (
    <img
      src={`/icons/${IconType.CHECKED}.svg`}
      alt='check box'
      height={height}
      width={width}
      className={['clickable', className].join(' ')}
    />
  )
}

export const BlankIcon: FC<Props> = ({
  height = 20,
  width = 20,
  className
}) => {
  return (
    <img
      src={`/icons/${IconType.BLANK}.svg`}
      alt='check box'
      height={height}
      width={width}
      className={['clickable', className].join(' ')}
    />
  )
}

export default CheckedIcon;