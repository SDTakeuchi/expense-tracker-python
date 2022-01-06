/*eslint-disable*/
import React from "react";
import DeleteIcon from "@material-ui/icons/Delete";
import IconButton from "@material-ui/core/IconButton";
// react components for routing our app without refresh
import { Link } from "react-router-dom";

// @material-ui/core components
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Tooltip from "@material-ui/core/Tooltip";

// @material-ui/icons
import { Home, Person, CloudDownload } from "@material-ui/icons";
import HomeIcon from '@material-ui/icons/Home';

// core components
import CustomDropdown from "components/CustomDropdown/CustomDropdown.js";
import Button from "components/CustomButtons/Button.js";

import styles from "assets/jss/material-kit-react/components/headerLinksStyle.js";

const useStyles = makeStyles(styles);

export default function HeaderLinks(props) {
  const classes = useStyles();
  return (
    <List className={classes.list}>
      <ListItem className={classes.listItem}>
        <IconButton aria-label="HomeLink">
          <HomeIcon />
        </IconButton>
      </ListItem>
      <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText=""
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={Person}
          dropdownList={[
            <Link to="/" className={classes.dropdownLink}>
              設定
            </Link>,
            <a
              href="https://creativetimofficial.github.io/material-kit-react/#/documentation?ref=mkr-navbar"
              target="_blank"
              className={classes.dropdownLink}
            >
              ログアウト
            </a>,
          ]}
        />
      </ListItem>
      {/* <ListItem className={classes.listItem}>
        <Button
          href="https://www.creative-tim.com/product/material-kit-react?ref=mkr-navbar"
          color="transparent"
          target="_blank"
          className={classes.navLink}
        >
          <CloudDownload className={classes.icons} /> Download
        </Button>
      </ListItem> */}
      <ListItem className={classes.listItem}>
        <Tooltip
          id="github-tooltip"
          title="Check my work on github"
          placement={window.innerWidth > 959 ? "top" : "left"}
          classes={{ tooltip: classes.tooltip }}
        >
          <Button
            color="transparent"
            href="https://github.com/SDTakeuchi"
            target="_blank"
            className={classes.navLink}
          >
            <i className={classes.socialIcons + " fab fa-home"} />
          </Button>
        </Tooltip>
      </ListItem>
      {/* <ListItem className={classes.listItem}>
        <Tooltip
          id="applemusic-tooltip"
          title="Apple Music"
          placement={window.innerWidth > 959 ? "top" : "left"}
          classes={{ tooltip: classes.tooltip }}
        >
          <Button
            color="transparent"
            href="https://music.apple.com/us/artist/douglas-takeuchi/1476471584"
            target="_blank"
            className={classes.navLink}
          >
            <i className={classes.socialIcons + " fab fa-apple"} />
          </Button>
        </Tooltip>
      </ListItem> */}
      {/* <ListItem className={classes.listItem}>
        <Tooltip
          id="user-settings-tooltip"
          title="User Settings"
          placement={window.innerWidth > 959 ? "top" : "left"}
          classes={{ tooltip: classes.tooltip }}
        >
          <Button
            color="transparent"
            href="https://open.spotify.com/artist/0TyGt3SnDcKDs5PsdHjfeT"
            target="_blank"
            className={classes.navLink}
          >
            <i className={classes.socialIcons + " fab fa-user-cogy"} />
          </Button>
        </Tooltip>
      </ListItem> */}
    </List>
  );
}
