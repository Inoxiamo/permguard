/**
 * Copyright 2024 Nitro Agility S.r.l.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Put your custom JS code here

const PYTHON_CODE = {
  before: `
def check_permissions(token: str, system: str, resource: str, action: str):
    payload = decode_jwt(token)
    roles: List[str] = payload.get("roles", [])
    for role in roles:
        role_permissions = get_permissions_for_role(role)
        if system in role_permissions:
            if resource in role_permissions[system]:
                if action in role_permissions[system][resource]:
                    # If all conditions match, permission is granted
                    return True
    # If no role grants permission, return False
    return False

has_permissions = check_permissions(token, system, "inventory", "view")`,
  after: `has_permissions = permguard.check(
    "uur::581616507495::iam:identity/google/pharmacist",
    "magicfarmacia-v0.0",
    "inventory",
    "view"
)`,
};

const GO_CODE = {
  before: `
func checkPermissions(token, system, resource, action string) bool {
    payload := decodeJWT(token)
    roles, ok := payload["roles"].([]string)
    if !ok {
        return false
    }

    for _, role := range roles {
        rolePermissions := getPermissionsForRole(role)
        if resources, systemFound := rolePermissions[resource]; systemFound {
            if actions, resourceFound := resources[system]; resourceFound {
                for _, allowedAction := range actions {
                    if strings.EqualFold(allowedAction, action) {
                        return true // Permission granted
                    }
                }
            }
        }
    }
    return false // No permission granted
}

hasPermissions := checkPermissions(token, system, "inventory", "view")`,
  after: `hasPermissions := permguard.Check(
    "uur::581616507495::iam:identity/google/pharmacist",
    "magicfarmacia-v0.0",
    "inventory",
    "view",
)`,
};

const handleScroll = () => {
  const header = document.querySelector("header");

  if (window.scrollY === 0) {
    header.classList.remove("header--blur");
  } else {
    header.classList.add("header--blur");
  }
};

handleScroll();

window.addEventListener("scroll", handleScroll);

// Language switch
let selectedLanguage = "go";
let isPermguard = false;

const handleSelectedLanguageChange = (element) => {
  const hasSeparator = window.innerWidth >= 768;

  const languageElements = document.querySelectorAll(".code__language");
  languageElements.forEach((el) => {
    el.classList.remove("code__language--active");
  });

  element.classList.add("code__language--active");
  selectedLanguage = element.getAttribute("data-language");

  const oldLanguage = selectedLanguage === "go" ? "python" : "go";

  let codeBoxes = [];

  if (hasSeparator) {
    codeBoxes = document.querySelectorAll("img-comparison-slider pre code");

    if (selectedLanguage === "go") {
      codeBoxes[0].innerHTML = GO_CODE.before;
      codeBoxes[1].innerHTML = GO_CODE.after;
    }

    if (selectedLanguage === "python") {
      codeBoxes[0].innerHTML = PYTHON_CODE.before;
      codeBoxes[1].innerHTML = PYTHON_CODE.after;
    }
  } else {
    codeBoxes = document.querySelectorAll(".code__img--small pre code");

    if (selectedLanguage === "go") {
      codeBoxes[0].innerHTML = GO_CODE[isPermguard ? "after" : "before"];
    }

    if (selectedLanguage === "python") {
      codeBoxes[0].innerHTML = PYTHON_CODE[isPermguard ? "after" : "before"];
    }
  }

  codeBoxes.forEach((codeBox) => {
    codeBox.classList.remove(`language-${oldLanguage}`);
    codeBox.classList.add(`language-${selectedLanguage}`);
    codeBox.removeAttribute("data-highlighted");
  });

  // eslint-disable-next-line no-undef
  hljs.highlightAll();
};

const languageElements = document.querySelectorAll(".code__language");
const switchInput = document.querySelector("#switchInput");

const toggleIsPermguard = () => {
  isPermguard = !isPermguard;

  const codeBoxes = document.querySelectorAll(".code__img--small pre code");

  if (selectedLanguage === "go") {
    codeBoxes[0].innerHTML = GO_CODE[isPermguard ? "after" : "before"];
  }

  if (selectedLanguage === "python") {
    codeBoxes[0].innerHTML = PYTHON_CODE[isPermguard ? "after" : "before"];
  }

  codeBoxes.forEach((codeBox) => {
    codeBox.removeAttribute("data-highlighted");
  });

  // eslint-disable-next-line no-undef
  hljs.highlightAll();
};

languageElements.forEach((el) => {
  el.addEventListener("click", () => handleSelectedLanguageChange(el));
});

if (switchInput) {
  switchInput.addEventListener("change", toggleIsPermguard);
}

// Detect forced dark mode
function detectForcedDarkMode() {
  if (
    !navigator.userAgent.match(/Samsung/i) ||
    !window.matchMedia ||
    window.matchMedia("(prefers-color-scheme:dark)").matches
  )
    return new Promise((resolve) => {
      resolve(false);
    });

  return new Promise((resolve) => {
    const ctx = document.createElement("canvas").getContext("2d");
    const svg = `
          <svg width="1" height="1" xmlns="http://www.w3.org/2000/svg">
              <rect width="1" height="1" fill="white"  />
          </svg>
      `;

    const blob = new Blob([svg], { type: "image/svg+xml" });
    const url = URL.createObjectURL(blob);

    const img = new Image();
    img.src = url;

    img.onload = () => {
      ctx.drawImage(img, 0, 0);
      const [r, g, b] = ctx.getImageData(0, 0, 1, 1).data;
      URL.revokeObjectURL(url); // Clean up the object URL to avoid memory leaks
      resolve((r & b & g) < 255);
    };

    img.onerror = () => {
      URL.revokeObjectURL(url); // Clean up on error as well
      resolve(false);
    };
  });
}

detectForcedDarkMode().then(function (isDarkModeForced) {
  if (isDarkModeForced) {
    document.querySelector("body").classList.add("forced-dark");
  }
});
