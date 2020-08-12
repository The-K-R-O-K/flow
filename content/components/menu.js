import React from "react";
import styled from "@emotion/styled";
import { NavItemTitle } from "gatsby-theme-flow";
import { size } from "polished";
import { theme } from "gatsby-theme-flow/src/colors";

export const MenuWrapper = styled.div({
  display: "grid",
  gridTemplateColumns: `repeat(auto-fill, minmax(270px, 1fr))`,
  gridGap: 24,
  paddingTop: 8,
});

const MenuItemWrapper = styled.div({
  display: "flex",
  li: {
    marginBottom: 0,
  },
});

const IconWrapper = styled.div({
  ...size(28),
  flexShrink: 0,
  marginRight: 16,
});

const TextWrapper = styled.div({
  color: theme.text1,
});

export function MenuItem({ icon, title, children, ...props }) {
  return (
    <MenuItemWrapper {...props}>
      <IconWrapper>{icon}</IconWrapper>
      <TextWrapper>
        <NavItemTitle>{title}</NavItemTitle>
        {children}
      </TextWrapper>
    </MenuItemWrapper>
  );
}
