package emoji

import (
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"
)

// These are the emojis the site can use.
// I used https://getemoji.com/ to get this list, but removed some emojis because these emojis didn't render properly in
// some browsers. There are of course more emojis, but the website has already enough emojis to choose from.
var emojis = "😀 😃 😄 😁 😆 😅 😂 🤣 🥲 \U0001F979 ☺️ 😊 😇 🙂 🙃 😉 😌 😍 🥰 😘 😗 😙 😚 😋 😛 😝 😜 🤪 🤨 🧐 🤓 😎 🥸 🤩 🥳 😏 😒 😞 😔 😟 😕 🙁 ☹️ 😣 😖 😫 😩 🥺 😢 😭 😮‍💨 😤 😠 😡 🤬 🤯 😳 🥵 🥶 😱 😨 😰 😥 😓 \U0001FAE3 🤗 \U0001FAE1 🤔 \U0001FAE2 🤭 🤫 🤥 😶 😶‍🌫️ 😐 😑 😬 \U0001FAE0 🙄 😯 😦 😧 😮 😲 🥱 😴 🤤 😪 😵 😵‍💫 \U0001FAE5 🤐 🥴 🤢 🤮 🤧 😷 🤒 🤕 🤑 🤠 😈 👿 👹 👺 🤡 💩 👻 💀 ☠️ 👽 👾 🤖 🎃 😺 😸 😹 😻 😼 😽 🙀 😿 😾 👋 🤚 🖐 ✋ 🖖 👌 🤌 🤏 ✌️ 🤞 \U0001FAF0 🤟 🤘 🤙 \U0001FAF5 \U0001FAF1 \U0001FAF2 \U0001FAF3 \U0001FAF4 👈 👉 👆 🖕 👇 ☝️ 👍 👎 ✊ 👊 🤛 🤜 👏 \U0001FAF6 🙌 👐 🤲 🤝 🙏 ✍️ 💅 🤳 💪 🦾 🦵 🦿 🦶 👣 👂 🦻 👃 🫀 🫁 🧠 🦷 🦴 👀 👁 👅 👄 \U0001FAE6 💋 🩸 👶 👧 🧒 👦 👩 🧑 👨 👩‍🦱 🧑‍🦱 👨‍🦱 👩‍🦰 🧑‍🦰 👨‍🦰 👱‍♀️ 👱 👱‍♂️ 👩‍🦳 🧑‍🦳 👨‍🦳 👩‍🦲 🧑‍🦲 👨‍🦲 🧔‍♀️ 🧔 🧔‍♂️ 👵 🧓 👴 👲 👳‍♀️ 👳 👳‍♂️ 🧕 👮‍♀️ 👮 👮‍♂️ 👷‍♀️ 👷 👷‍♂️ 💂‍♀️ 💂 💂‍♂️ 🕵️‍♀️ 🕵️ 🕵️‍♂️ 👩‍⚕️ 🧑‍⚕️ 👨‍⚕️ 👩‍🌾 🧑‍🌾 👨‍🌾 👩‍🍳 🧑‍🍳 👨‍🍳 👩‍🎓 🧑‍🎓 👨‍🎓 👩‍🎤 🧑‍🎤 👨‍🎤 👩‍🏫 🧑‍🏫 👨‍🏫 👩‍🏭 🧑‍🏭 👨‍🏭 👩‍💻 🧑‍💻 👨‍💻 👩‍💼 🧑‍💼 👨‍💼 👩‍🔧 🧑‍🔧 👨‍🔧 👩‍🔬 🧑‍🔬 👨‍🔬 👩‍🎨 🧑‍🎨 👨‍🎨 👩‍🚒 🧑‍🚒 👨‍🚒 👩‍✈️ 🧑‍✈️ 👨‍✈️ 👩‍🚀 🧑‍🚀 👨‍🚀 👩‍⚖️ 🧑‍⚖️ 👨‍⚖️ 👰‍♀️ 👰 👰‍♂️ 🤵‍♀️ 🤵 🤵‍♂️ 👸 \U0001FAC5 🤴 🥷 🦸‍♀️ 🦸 🦸‍♂️ 🦹‍♀️ 🦹 🦹‍♂️ 🤶 🧑‍🎄 🎅 🧙‍♀️ 🧙 🧙‍♂️ 🧝‍♀️ 🧝 🧝‍♂️ 🧛‍♀️ 🧛 🧛‍♂️ 🧟‍♀️ 🧟 🧟‍♂️ 🧞‍♀️ 🧞 🧞‍♂️ 🧜‍♀️ 🧜 🧜‍♂️ 🧚‍♀️ 🧚 🧚‍♂️ \U0001F9CC 👼 🤰 \U0001FAC4 \U0001FAC3 🤱 👩‍🍼 🧑‍🍼 👨‍🍼 🙇‍♀️ 🙇 🙇‍♂️ 💁‍♀️ 💁 💁‍♂️ 🙅‍♀️ 🙅 🙅‍♂️ 🙆‍♀️ 🙆 🙆‍♂️ 🙋‍♀️ 🙋 🙋‍♂️ 🧏‍♀️ 🧏 🧏‍♂️ 🤦‍♀️ 🤦 🤦‍♂️ 🤷‍♀️ 🤷 🤷‍♂️ 🙎‍♀️ 🙎 🙎‍♂️ 🙍‍♀️ 🙍 🙍‍♂️ 💇‍♀️ 💇 💇‍♂️ 💆‍♀️ 💆 💆‍♂️ 🧖‍♀️ 🧖 🧖‍♂️ 💅 🤳 💃 🕺 👯‍♀️ 👯 👯‍♂️ 🕴 👩‍🦽 🧑‍🦽 👨‍🦽 👩‍🦼 🧑‍🦼 👨‍🦼 🚶‍♀️ 🚶 🚶‍♂️ 👩‍🦯 🧑‍🦯 👨‍🦯 🧎‍♀️ 🧎 🧎‍♂️ 🏃‍♀️ 🏃 🏃‍♂️ 🧍‍♀️ 🧍 🧍‍♂️ 👭 🧑‍🤝‍🧑 👬 👫 👩‍❤️‍👩 💑 👨‍❤️‍👨 👩‍❤️‍👨 👩‍❤️‍💋‍👩 💏 👨‍❤️‍💋‍👨 👩‍❤️‍💋‍👨 👪 👨‍👩‍👦 👨‍👩‍👧 👨‍👩‍👧‍👦 👨‍👩‍👦‍👦 👨‍👩‍👧‍👧 👨‍👨‍👦 👨‍👨‍👧 👨‍👨‍👧‍👦 👨‍👨‍👦‍👦 👨‍👨‍👧‍👧 👩‍👩‍👦 👩‍👩‍👧 👩‍👩‍👧‍👦 👩‍👩‍👦‍👦 👩‍👩‍👧‍👧 👨‍👦 👨‍👦‍👦 👨‍👧 👨‍👧‍👦 👨‍👧‍👧 👩‍👦 👩‍👦‍👦 👩‍👧 👩‍👧‍👦 👩‍👧‍👧 🗣 👤 👥 🫂 🧳 🌂 ☂️ 🧵 🪡 🪢 🧶 👓 🕶 🥽 🥼 🦺 👔 👕 👖 🧣 🧤 🧥 🧦 👗 👘 🥻 🩴 🩱 🩲 🩳 👙 👚 👛 👜 👝 🎒 👞 👟 🥾 🥿 👠 👡 🩰 👢 👑 👒 🎩 🎓 🧢 ⛑ 🪖 💄 💍 💼 👋🏻 🤚🏻 🖐🏻 ✋🏻 🖖🏻 👌🏻 🤌🏻 🤏🏻 ✌🏻 🤞🏻 \U0001FAF0🏻 🤟🏻 🤘🏻 🤙🏻 \U0001FAF5🏻 \U0001FAF1🏻 \U0001FAF2🏻 \U0001FAF3🏻 \U0001FAF4🏻 👈🏻 👉🏻 👆🏻 🖕🏻 👇🏻 ☝🏻 👍🏻 👎🏻 ✊🏻 👊🏻 🤛🏻 🤜🏻 👏🏻 \U0001FAF6🏻 🙌🏻 👐🏻 🤲🏻 🙏🏻 ✍🏻 💅🏻 🤳🏻 💪🏻 🦵🏻 🦶🏻 👂🏻 🦻🏻 👃🏻 👶🏻 👧🏻 🧒🏻 👦🏻 👩🏻 🧑🏻 👨🏻 👩🏻‍🦱 🧑🏻‍🦱 👨🏻‍🦱 👩🏻‍🦰 🧑🏻‍🦰 👨🏻‍🦰 👱🏻‍♀️ 👱🏻 👱🏻‍♂️ 👩🏻‍🦳 🧑🏻‍🦳 👨🏻‍🦳 👩🏻‍🦲 🧑🏻‍🦲 👨🏻‍🦲 🧔🏻‍♀️ 🧔🏻 🧔🏻‍♂️ 👵🏻 🧓🏻 👴🏻 👲🏻 👳🏻‍♀️ 👳🏻 👳🏻‍♂️ 🧕🏻 👮🏻‍♀️ 👮🏻 👮🏻‍♂️ 👷🏻‍♀️ 👷🏻 👷🏻‍♂️ 💂🏻‍♀️ 💂🏻 💂🏻‍♂️ 🕵🏻‍♀️ 🕵🏻 🕵🏻‍♂️ 👩🏻‍⚕️ 🧑🏻‍⚕️ 👨🏻‍⚕️ 👩🏻‍🌾 🧑🏻‍🌾 👨🏻‍🌾 👩🏻‍🍳 🧑🏻‍🍳 👨🏻‍🍳 👩🏻‍🎓 🧑🏻‍🎓 👨🏻‍🎓 👩🏻‍🎤 🧑🏻‍🎤 👨🏻‍🎤 👩🏻‍🏫 🧑🏻‍🏫 👨🏻‍🏫 👩🏻‍🏭 🧑🏻‍🏭 👨🏻‍🏭 👩🏻‍💻 🧑🏻‍💻 👨🏻‍💻 👩🏻‍💼 🧑🏻‍💼 👨🏻‍💼 👩🏻‍🔧 🧑🏻‍🔧 👨🏻‍🔧 👩🏻‍🔬 🧑🏻‍🔬 👨🏻‍🔬 👩🏻‍🎨 🧑🏻‍🎨 👨🏻‍🎨 👩🏻‍🚒 🧑🏻‍🚒 👨🏻‍🚒 👩🏻‍✈️ 🧑🏻‍✈️ 👨🏻‍✈️ 👩🏻‍🚀 🧑🏻‍🚀 👨🏻‍🚀 👩🏻‍⚖️ 🧑🏻‍⚖️ 👨🏻‍⚖️ 👰🏻‍♀️ 👰🏻 👰🏻‍♂️ 🤵🏻‍♀️ 🤵🏻 🤵🏻‍♂️ 👸🏻 \U0001FAC5🏻 🤴🏻 🥷🏻 🦸🏻‍♀️ 🦸🏻 🦸🏻‍♂️ 🦹🏻‍♀️ 🦹🏻 🦹🏻‍♂️ 🤶🏻 🧑🏻‍🎄 🎅🏻 🧙🏻‍♀️ 🧙🏻 🧙🏻‍♂️ 🧝🏻‍♀️ 🧝🏻 🧝🏻‍♂️ 🧛🏻‍♀️ 🧛🏻 🧛🏻‍♂️ 🧜🏻‍♀️ 🧜🏻 🧜🏻‍♂️ 🧚🏻‍♀️ 🧚🏻 🧚🏻‍♂️ 👼🏻 🤰🏻 \U0001FAC4🏻 \U0001FAC3🏻 🤱🏻 👩🏻‍🍼 🧑🏻‍🍼 👨🏻‍🍼 🙇🏻‍♀️ 🙇🏻 🙇🏻‍♂️ 💁🏻‍♀️ 💁🏻 💁🏻‍♂️ 🙅🏻‍♀️ 🙅🏻 🙅🏻‍♂️ 🙆🏻‍♀️ 🙆🏻 🙆🏻‍♂️ 🙋🏻‍♀️ 🙋🏻 🙋🏻‍♂️ 🧏🏻‍♀️ 🧏🏻 🧏🏻‍♂️ 🤦🏻‍♀️ 🤦🏻 🤦🏻‍♂️ 🤷🏻‍♀️ 🤷🏻 🤷🏻‍♂️ 🙎🏻‍♀️ 🙎🏻 🙎🏻‍♂️ 🙍🏻‍♀️ 🙍🏻 🙍🏻‍♂️ 💇🏻‍♀️ 💇🏻 💇🏻‍♂️ 💆🏻‍♀️ 💆🏻 💆🏻‍♂️ 🧖🏻‍♀️ 🧖🏻 🧖🏻‍♂️ 💃🏻 🕺🏻 🕴🏻 👩🏻‍🦽 🧑🏻‍🦽 👨🏻‍🦽 👩🏻‍🦼 🧑🏻‍🦼 👨🏻‍🦼 🚶🏻‍♀️ 🚶🏻 🚶🏻‍♂️ 👩🏻‍🦯 🧑🏻‍🦯 👨🏻‍🦯 🧎🏻‍♀️ 🧎🏻 🧎🏻‍♂️ 🏃🏻‍♀️ 🏃🏻 🏃🏻‍♂️ 🧍🏻‍♀️ 🧍🏻 🧍🏻‍♂️ 👭🏻 🧑🏻‍🤝‍🧑🏻 👬🏻 👫🏻 🧗🏻‍♀️ 🧗🏻 🧗🏻‍♂️ 🏇🏻 🏂🏻 🏌🏻‍♀️ 🏌🏻 🏌🏻‍♂️ 🏄🏻‍♀️ 🏄🏻 🏄🏻‍♂️ 🚣🏻‍♀️ 🚣🏻 🚣🏻‍♂️ 🏊🏻‍♀️ 🏊🏻 🏊🏻‍♂️ ⛹🏻‍♀️ ⛹🏻 ⛹🏻‍♂️ 🏋🏻‍♀️ 🏋🏻 🏋🏻‍♂️ 🚴🏻‍♀️ 🚴🏻 🚴🏻‍♂️ 🚵🏻‍♀️ 🚵🏻 🚵🏻‍♂️ 🤸🏻‍♀️ 🤸🏻 🤸🏻‍♂️ 🤽🏻‍♀️ 🤽🏻 🤽🏻‍♂️ 🤾🏻‍♀️ 🤾🏻 🤾🏻‍♂️ 🤹🏻‍♀️ 🤹🏻 🤹🏻‍♂️ 🧘🏻‍♀️ 🧘🏻 🧘🏻‍♂️ 🛀🏻 🛌🏻 👋🏼 🤚🏼 🖐🏼 ✋🏼 🖖🏼 👌🏼 🤌🏼 🤏🏼 ✌🏼 🤞🏼 \U0001FAF0🏼 🤟🏼 🤘🏼 🤙🏼 \U0001FAF5🏼 \U0001FAF1🏼 \U0001FAF2🏼 \U0001FAF3🏼 \U0001FAF4🏼 👈🏼 👉🏼 👆🏼 🖕🏼 👇🏼 ☝🏼 👍🏼 👎🏼 ✊🏼 👊🏼 🤛🏼 🤜🏼 👏🏼 \U0001FAF6🏼 🙌🏼 👐🏼 🤲🏼 🙏🏼 ✍🏼 💅🏼 🤳🏼 💪🏼 🦵🏼 🦶🏼 👂🏼 🦻🏼 👃🏼 👶🏼 👧🏼 🧒🏼 👦🏼 👩🏼 🧑🏼 👨🏼 👩🏼‍🦱 🧑🏼‍🦱 👨🏼‍🦱 👩🏼‍🦰 🧑🏼‍🦰 👨🏼‍🦰 👱🏼‍♀️ 👱🏼 👱🏼‍♂️ 👩🏼‍🦳 🧑🏼‍🦳 👨🏼‍🦳 👩🏼‍🦲 🧑🏼‍🦲 👨🏼‍🦲 🧔🏼‍♀️ 🧔🏼 🧔🏼‍♂️ 👵🏼 🧓🏼 👴🏼 👲🏼 👳🏼‍♀️ 👳🏼 👳🏼‍♂️ 🧕🏼 👮🏼‍♀️ 👮🏼 👮🏼‍♂️ 👷🏼‍♀️ 👷🏼 👷🏼‍♂️ 💂🏼‍♀️ 💂🏼 💂🏼‍♂️ 🕵🏼‍♀️ 🕵🏼 🕵🏼‍♂️ 👩🏼‍⚕️ 🧑🏼‍⚕️ 👨🏼‍⚕️ 👩🏼‍🌾 🧑🏼‍🌾 👨🏼‍🌾 👩🏼‍🍳 🧑🏼‍🍳 👨🏼‍🍳 👩🏼‍🎓 🧑🏼‍🎓 👨🏼‍🎓 👩🏼‍🎤 🧑🏼‍🎤 👨🏼‍🎤 👩🏼‍🏫 🧑🏼‍🏫 👨🏼‍🏫 👩🏼‍🏭 🧑🏼‍🏭 👨🏼‍🏭 👩🏼‍💻 🧑🏼‍💻 👨🏼‍💻 👩🏼‍💼 🧑🏼‍💼 👨🏼‍💼 👩🏼‍🔧 🧑🏼‍🔧 👨🏼‍🔧 👩🏼‍🔬 🧑🏼‍🔬 👨🏼‍🔬 👩🏼‍🎨 🧑🏼‍🎨 👨🏼‍🎨 👩🏼‍🚒 🧑🏼‍🚒 👨🏼‍🚒 👩🏼‍✈️ 🧑🏼‍✈️ 👨🏼‍✈️ 👩🏼‍🚀 🧑🏼‍🚀 👨🏼‍🚀 👩🏼‍⚖️ 🧑🏼‍⚖️ 👨🏼‍⚖️ 👰🏼‍♀️ 👰🏼 👰🏼‍♂️ 🤵🏼‍♀️ 🤵🏼 🤵🏼‍♂️ 👸🏼 \U0001FAC5🏼 🤴🏼 🥷🏼 🦸🏼‍♀️ 🦸🏼 🦸🏼‍♂️ 🦹🏼‍♀️ 🦹🏼 🦹🏼‍♂️ 🤶🏼 🧑🏼‍🎄 🎅🏼 🧙🏼‍♀️ 🧙🏼 🧙🏼‍♂️ 🧝🏼‍♀️ 🧝🏼 🧝🏼‍♂️ 🧛🏼‍♀️ 🧛🏼 🧛🏼‍♂️ 🧜🏼‍♀️ 🧜🏼 🧜🏼‍♂️ 🧚🏼‍♀️ 🧚🏼 🧚🏼‍♂️ 👼🏼 🤰🏼 \U0001FAC4🏼 \U0001FAC3🏼 🤱🏼 👩🏼‍🍼 🧑🏼‍🍼 👨🏼‍🍼 🙇🏼‍♀️ 🙇🏼 🙇🏼‍♂️ 💁🏼‍♀️ 💁🏼 💁🏼‍♂️ 🙅🏼‍♀️ 🙅🏼 🙅🏼‍♂️ 🙆🏼‍♀️ 🙆🏼 🙆🏼‍♂️ 🙋🏼‍♀️ 🙋🏼 🙋🏼‍♂️ 🧏🏼‍♀️ 🧏🏼 🧏🏼‍♂️ 🤦🏼‍♀️ 🤦🏼 🤦🏼‍♂️ 🤷🏼‍♀️ 🤷🏼 🤷🏼‍♂️ 🙎🏼‍♀️ 🙎🏼 🙎🏼‍♂️ 🙍🏼‍♀️ 🙍🏼 🙍🏼‍♂️ 💇🏼‍♀️ 💇🏼 💇🏼‍♂️ 💆🏼‍♀️ 💆🏼 💆🏼‍♂️ 🧖🏼‍♀️ 🧖🏼 🧖🏼‍♂️ 💃🏼 🕺🏼 🕴🏼 👩🏼‍🦽 🧑🏼‍🦽 👨🏼‍🦽 👩🏼‍🦼 🧑🏼‍🦼 👨🏼‍🦼 🚶🏼‍♀️ 🚶🏼 🚶🏼‍♂️ 👩🏼‍🦯 🧑🏼‍🦯 👨🏼‍🦯 🧎🏼‍♀️ 🧎🏼 🧎🏼‍♂️ 🏃🏼‍♀️ 🏃🏼 🏃🏼‍♂️ 🧍🏼‍♀️ 🧍🏼 🧍🏼‍♂️ 👭🏼 🧑🏼‍🤝‍🧑🏼 👬🏼 👫🏼 🧗🏼‍♀️ 🧗🏼 🧗🏼‍♂️ 🏇🏼 🏂🏼 🏌🏼‍♀️ 🏌🏼 🏌🏼‍♂️ 🏄🏼‍♀️ 🏄🏼 🏄🏼‍♂️ 🚣🏼‍♀️ 🚣🏼 🚣🏼‍♂️ 🏊🏼‍♀️ 🏊🏼 🏊🏼‍♂️ ⛹🏼‍♀️ ⛹🏼 ⛹🏼‍♂️ 🏋🏼‍♀️ 🏋🏼 🏋🏼‍♂️ 🚴🏼‍♀️ 🚴🏼 🚴🏼‍♂️ 🚵🏼‍♀️ 🚵🏼 🚵🏼‍♂️ 🤸🏼‍♀️ 🤸🏼 🤸🏼‍♂️ 🤽🏼‍♀️ 🤽🏼 🤽🏼‍♂️ 🤾🏼‍♀️ 🤾🏼 🤾🏼‍♂️ 🤹🏼‍♀️ 🤹🏼 🤹🏼‍♂️ 🧘🏼‍♀️ 🧘🏼 🧘🏼‍♂️ 🛀🏼 🛌🏼 👋🏽 🤚🏽 🖐🏽 ✋🏽 🖖🏽 👌🏽 🤌🏽 🤏🏽 ✌🏽 🤞🏽 \U0001FAF0🏽 🤟🏽 🤘🏽 🤙🏽 \U0001FAF5🏽 \U0001FAF1🏽 \U0001FAF2🏽 \U0001FAF3🏽 \U0001FAF4🏽 👈🏽 👉🏽 👆🏽 🖕🏽 👇🏽 ☝🏽 👍🏽 👎🏽 ✊🏽 👊🏽 🤛🏽 🤜🏽 👏🏽 \U0001FAF6🏽 🙌🏽 👐🏽 🤲🏽 🙏🏽 ✍🏽 💅🏽 🤳🏽 💪🏽 🦵🏽 🦶🏽 👂🏽 🦻🏽 👃🏽 👶🏽 👧🏽 🧒🏽 👦🏽 👩🏽 🧑🏽 👨🏽 👩🏽‍🦱 🧑🏽‍🦱 👨🏽‍🦱 👩🏽‍🦰 🧑🏽‍🦰 👨🏽‍🦰 👱🏽‍♀️ 👱🏽 👱🏽‍♂️ 👩🏽‍🦳 🧑🏽‍🦳 👨🏽‍🦳 👩🏽‍🦲 🧑🏽‍🦲 👨🏽‍🦲 🧔🏽‍♀️ 🧔🏽 🧔🏽‍♂️ 👵🏽 🧓🏽 👴🏽 👲🏽 👳🏽‍♀️ 👳🏽 👳🏽‍♂️ 🧕🏽 👮🏽‍♀️ 👮🏽 👮🏽‍♂️ 👷🏽‍♀️ 👷🏽 👷🏽‍♂️ 💂🏽‍♀️ 💂🏽 💂🏽‍♂️ 🕵🏽‍♀️ 🕵🏽 🕵🏽‍♂️ 👩🏽‍⚕️ 🧑🏽‍⚕️ 👨🏽‍⚕️ 👩🏽‍🌾 🧑🏽‍🌾 👨🏽‍🌾 👩🏽‍🍳 🧑🏽‍🍳 👨🏽‍🍳 👩🏽‍🎓 🧑🏽‍🎓 👨🏽‍🎓 👩🏽‍🎤 🧑🏽‍🎤 👨🏽‍🎤 👩🏽‍🏫 🧑🏽‍🏫 👨🏽‍🏫 👩🏽‍🏭 🧑🏽‍🏭 👨🏽‍🏭 👩🏽‍💻 🧑🏽‍💻 👨🏽‍💻 👩🏽‍💼 🧑🏽‍💼 👨🏽‍💼 👩🏽‍🔧 🧑🏽‍🔧 👨🏽‍🔧 👩🏽‍🔬 🧑🏽‍🔬 👨🏽‍🔬 👩🏽‍🎨 🧑🏽‍🎨 👨🏽‍🎨 👩🏽‍🚒 🧑🏽‍🚒 👨🏽‍🚒 👩🏽‍✈️ 🧑🏽‍✈️ 👨🏽‍✈️ 👩🏽‍🚀 🧑🏽‍🚀 👨🏽‍🚀 👩🏽‍⚖️ 🧑🏽‍⚖️ 👨🏽‍⚖️ 👰🏽‍♀️ 👰🏽 👰🏽‍♂️ 🤵🏽‍♀️ 🤵🏽 🤵🏽‍♂️ 👸🏽 \U0001FAC5🏽 🤴🏽 🥷🏽 🦸🏽‍♀️ 🦸🏽 🦸🏽‍♂️ 🦹🏽‍♀️ 🦹🏽 🦹🏽‍♂️ 🤶🏽 🧑🏽‍🎄 🎅🏽 🧙🏽‍♀️ 🧙🏽 🧙🏽‍♂️ 🧝🏽‍♀️ 🧝🏽 🧝🏽‍♂️ 🧛🏽‍♀️ 🧛🏽 🧛🏽‍♂️ 🧜🏽‍♀️ 🧜🏽 🧜🏽‍♂️ 🧚🏽‍♀️ 🧚🏽 🧚🏽‍♂️ 👼🏽 🤰🏽 \U0001FAC4🏽 \U0001FAC3🏽 🤱🏽 👩🏽‍🍼 🧑🏽‍🍼 👨🏽‍🍼 🙇🏽‍♀️ 🙇🏽 🙇🏽‍♂️ 💁🏽‍♀️ 💁🏽 💁🏽‍♂️ 🙅🏽‍♀️ 🙅🏽 🙅🏽‍♂️ 🙆🏽‍♀️ 🙆🏽 🙆🏽‍♂️ 🙋🏽‍♀️ 🙋🏽 🙋🏽‍♂️ 🧏🏽‍♀️ 🧏🏽 🧏🏽‍♂️ 🤦🏽‍♀️ 🤦🏽 🤦🏽‍♂️ 🤷🏽‍♀️ 🤷🏽 🤷🏽‍♂️ 🙎🏽‍♀️ 🙎🏽 🙎🏽‍♂️ 🙍🏽‍♀️ 🙍🏽 🙍🏽‍♂️ 💇🏽‍♀️ 💇🏽 💇🏽‍♂️ 💆🏽‍♀️ 💆🏽 💆🏽‍♂️ 🧖🏽‍♀️ 🧖🏽 🧖🏽‍♂️ 💃🏽 🕺🏽 🕴🏽 👩🏽‍🦽 🧑🏽‍🦽 👨🏽‍🦽 👩🏽‍🦼 🧑🏽‍🦼 👨🏽‍🦼 🚶🏽‍♀️ 🚶🏽 🚶🏽‍♂️ 👩🏽‍🦯 🧑🏽‍🦯 👨🏽‍🦯 🧎🏽‍♀️ 🧎🏽 🧎🏽‍♂️ 🏃🏽‍♀️ 🏃🏽 🏃🏽‍♂️ 🧍🏽‍♀️ 🧍🏽 🧍🏽‍♂️ 👭🏽 🧑🏽‍🤝‍🧑🏽 👬🏽 👫🏽 🧗🏽‍♀️ 🧗🏽 🧗🏽‍♂️ 🏇🏽 🏂🏽 🏌🏽‍♀️ 🏌🏽 🏌🏽‍♂️ 🏄🏽‍♀️ 🏄🏽 🏄🏽‍♂️ 🚣🏽‍♀️ 🚣🏽 🚣🏽‍♂️ 🏊🏽‍♀️ 🏊🏽 🏊🏽‍♂️ ⛹🏽‍♀️ ⛹🏽 ⛹🏽‍♂️ 🏋🏽‍♀️ 🏋🏽 🏋🏽‍♂️ 🚴🏽‍♀️ 🚴🏽 🚴🏽‍♂️ 🚵🏽‍♀️ 🚵🏽 🚵🏽‍♂️ 🤸🏽‍♀️ 🤸🏽 🤸🏽‍♂️ 🤽🏽‍♀️ 🤽🏽 🤽🏽‍♂️ 🤾🏽‍♀️ 🤾🏽 🤾🏽‍♂️ 🤹🏽‍♀️ 🤹🏽 🤹🏽‍♂️ 🧘🏽‍♀️ 🧘🏽 🧘🏽‍♂️ 🛀🏽 🛌🏽 👋🏾 🤚🏾 🖐🏾 ✋🏾 🖖🏾 👌🏾 🤌🏾 🤏🏾 ✌🏾 🤞🏾 \U0001FAF0🏾 🤟🏾 🤘🏾 🤙🏾 \U0001FAF5🏾 \U0001FAF1🏾 \U0001FAF2🏾 \U0001FAF3🏾 \U0001FAF4🏾 👈🏾 👉🏾 👆🏾 🖕🏾 👇🏾 ☝🏾 👍🏾 👎🏾 ✊🏾 👊🏾 🤛🏾 🤜🏾 👏🏾 \U0001FAF6🏾 🙌🏾 👐🏾 🤲🏾 🙏🏾 ✍🏾 💅🏾 🤳🏾 💪🏾 🦵🏾 🦶🏾 👂🏾 🦻🏾 👃🏾 👶🏾 👧🏾 🧒🏾 👦🏾 👩🏾 🧑🏾 👨🏾 👩🏾‍🦱 🧑🏾‍🦱 👨🏾‍🦱 👩🏾‍🦰 🧑🏾‍🦰 👨🏾‍🦰 👱🏾‍♀️ 👱🏾 👱🏾‍♂️ 👩🏾‍🦳 🧑🏾‍🦳 👨🏾‍🦳 👩🏾‍🦲 🧑🏾‍🦲 👨🏾‍🦲 🧔🏾‍♀️ 🧔🏾 🧔🏾‍♂️ 👵🏾 🧓🏾 👴🏾 👲🏾 👳🏾‍♀️ 👳🏾 👳🏾‍♂️ 🧕🏾 👮🏾‍♀️ 👮🏾 👮🏾‍♂️ 👷🏾‍♀️ 👷🏾 👷🏾‍♂️ 💂🏾‍♀️ 💂🏾 💂🏾‍♂️ 🕵🏾‍♀️ 🕵🏾 🕵🏾‍♂️ 👩🏾‍⚕️ 🧑🏾‍⚕️ 👨🏾‍⚕️ 👩🏾‍🌾 🧑🏾‍🌾 👨🏾‍🌾 👩🏾‍🍳 🧑🏾‍🍳 👨🏾‍🍳 👩🏾‍🎓 🧑🏾‍🎓 👨🏾‍🎓 👩🏾‍🎤 🧑🏾‍🎤 👨🏾‍🎤 👩🏾‍🏫 🧑🏾‍🏫 👨🏾‍🏫 👩🏾‍🏭 🧑🏾‍🏭 👨🏾‍🏭 👩🏾‍💻 🧑🏾‍💻 👨🏾‍💻 👩🏾‍💼 🧑🏾‍💼 👨🏾‍💼 👩🏾‍🔧 🧑🏾‍🔧 👨🏾‍🔧 👩🏾‍🔬 🧑🏾‍🔬 👨🏾‍🔬 👩🏾‍🎨 🧑🏾‍🎨 👨🏾‍🎨 👩🏾‍🚒 🧑🏾‍🚒 👨🏾‍🚒 👩🏾‍✈️ 🧑🏾‍✈️ 👨🏾‍✈️ 👩🏾‍🚀 🧑🏾‍🚀 👨🏾‍🚀 👩🏾‍⚖️ 🧑🏾‍⚖️ 👨🏾‍⚖️ 👰🏾‍♀️ 👰🏾 👰🏾‍♂️ 🤵🏾‍♀️ 🤵🏾 🤵🏾‍♂️ 👸🏾 \U0001FAC5🏾 🤴🏾 🥷🏾 🦸🏾‍♀️ 🦸🏾 🦸🏾‍♂️ 🦹🏾‍♀️ 🦹🏾 🦹🏾‍♂️ 🤶🏾 🧑🏾‍🎄 🎅🏾 🧙🏾‍♀️ 🧙🏾 🧙🏾‍♂️ 🧝🏾‍♀️ 🧝🏾 🧝🏾‍♂️ 🧛🏾‍♀️ 🧛🏾 🧛🏾‍♂️ 🧜🏾‍♀️ 🧜🏾 🧜🏾‍♂️ 🧚🏾‍♀️ 🧚🏾 🧚🏾‍♂️ 👼🏾 🤰🏾 \U0001FAC4🏾 \U0001FAC3🏾 🤱🏾 👩🏾‍🍼 🧑🏾‍🍼 👨🏾‍🍼 🙇🏾‍♀️ 🙇🏾 🙇🏾‍♂️ 💁🏾‍♀️ 💁🏾 💁🏾‍♂️ 🙅🏾‍♀️ 🙅🏾 🙅🏾‍♂️ 🙆🏾‍♀️ 🙆🏾 🙆🏾‍♂️ 🙋🏾‍♀️ 🙋🏾 🙋🏾‍♂️ 🧏🏾‍♀️ 🧏🏾 🧏🏾‍♂️ 🤦🏾‍♀️ 🤦🏾 🤦🏾‍♂️ 🤷🏾‍♀️ 🤷🏾 🤷🏾‍♂️ 🙎🏾‍♀️ 🙎🏾 🙎🏾‍♂️ 🙍🏾‍♀️ 🙍🏾 🙍🏾‍♂️ 💇🏾‍♀️ 💇🏾 💇🏾‍♂️ 💆🏾‍♀️ 💆🏾 💆🏾‍♂️ 🧖🏾‍♀️ 🧖🏾 🧖🏾‍♂️ 💃🏾 🕺🏾 🕴🏿 👩🏾‍🦽 🧑🏾‍🦽 👨🏾‍🦽 👩🏾‍🦼 🧑🏾‍🦼 👨🏾‍🦼 🚶🏾‍♀️ 🚶🏾 🚶🏾‍♂️ 👩🏾‍🦯 🧑🏾‍🦯 👨🏾‍🦯 🧎🏾‍♀️ 🧎🏾 🧎🏾‍♂️ 🏃🏾‍♀️ 🏃🏾 🏃🏾‍♂️ 🧍🏾‍♀️ 🧍🏾 🧍🏾‍♂️ 👭🏾 🧑🏾‍🤝‍🧑🏾 👬🏾 👫🏾 🧗🏾‍♀️ 🧗🏾 🧗🏾‍♂️ 🏇🏾 🏂🏾 🏌🏾‍♀️ 🏌🏾 🏌🏾‍♂️ 🏄🏾‍♀️ 🏄🏾 🏄🏾‍♂️ 🚣🏾‍♀️ 🚣🏾 🚣🏾‍♂️ 🏊🏾‍♀️ 🏊🏾 🏊🏾‍♂️ ⛹🏾‍♀️ ⛹🏾 ⛹🏾‍♂️ 🏋🏾‍♀️ 🏋🏾 🏋🏾‍♂️ 🚴🏾‍♀️ 🚴🏾 🚴🏾‍♂️ 🚵🏾‍♀️ 🚵🏾 🚵🏾‍♂️ 🤸🏾‍♀️ 🤸🏾 🤸🏾‍♂️ 🤽🏾‍♀️ 🤽🏾 🤽🏾‍♂️ 🤾🏾‍♀️ 🤾🏾 🤾🏾‍♂️ 🤹🏾‍♀️ 🤹🏾 🤹🏾‍♂️ 🧘🏾‍♀️ 🧘🏾 🧘🏾‍♂️ 🛀🏾 🛌🏾 👋🏿 🤚🏿 🖐🏿 ✋🏿 🖖🏿 👌🏿 🤌🏿 🤏🏿 ✌🏿 🤞🏿 \U0001FAF0🏿 🤟🏿 🤘🏿 🤙🏿 \U0001FAF5🏿 \U0001FAF1🏿 \U0001FAF2🏿 \U0001FAF3🏿 \U0001FAF4🏿 👈🏿 👉🏿 👆🏿 🖕🏿 👇🏿 ☝🏿 👍🏿 👎🏿 ✊🏿 👊🏿 🤛🏿 🤜🏿 👏🏿 \U0001FAF6🏿 🙌🏿 👐🏿 🤲🏿 🙏🏿 ✍🏿 💅🏿 🤳🏿 💪🏿 🦵🏿 🦶🏿 👂🏿 🦻🏿 👃🏿 👶🏿 👧🏿 🧒🏿 👦🏿 👩🏿 🧑🏿 👨🏿 👩🏿‍🦱 🧑🏿‍🦱 👨🏿‍🦱 👩🏿‍🦰 🧑🏿‍🦰 👨🏿‍🦰 👱🏿‍♀️ 👱🏿 👱🏿‍♂️ 👩🏿‍🦳 🧑🏿‍🦳 👨🏿‍🦳 👩🏿‍🦲 🧑🏿‍🦲 👨🏿‍🦲 🧔🏿‍♀️ 🧔🏿 🧔🏿‍♂️ 👵🏿 🧓🏿 👴🏿 👲🏿 👳🏿‍♀️ 👳🏿 👳🏿‍♂️ 🧕🏿 👮🏿‍♀️ 👮🏿 👮🏿‍♂️ 👷🏿‍♀️ 👷🏿 👷🏿‍♂️ 💂🏿‍♀️ 💂🏿 💂🏿‍♂️ 🕵🏿‍♀️ 🕵🏿 🕵🏿‍♂️ 👩🏿‍⚕️ 🧑🏿‍⚕️ 👨🏿‍⚕️ 👩🏿‍🌾 🧑🏿‍🌾 👨🏿‍🌾 👩🏿‍🍳 🧑🏿‍🍳 👨🏿‍🍳 👩🏿‍🎓 🧑🏿‍🎓 👨🏿‍🎓 👩🏿‍🎤 🧑🏿‍🎤 👨🏿‍🎤 👩🏿‍🏫 🧑🏿‍🏫 👨🏿‍🏫 👩🏿‍🏭 🧑🏿‍🏭 👨🏿‍🏭 👩🏿‍💻 🧑🏿‍💻 👨🏿‍💻 👩🏿‍💼 🧑🏿‍💼 👨🏿‍💼 👩🏿‍🔧 🧑🏿‍🔧 👨🏿‍🔧 👩🏿‍🔬 🧑🏿‍🔬 👨🏿‍🔬 👩🏿‍🎨 🧑🏿‍🎨 👨🏿‍🎨 👩🏿‍🚒 🧑🏿‍🚒 👨🏿‍🚒 👩🏿‍✈️ 🧑🏿‍✈️ 👨🏿‍✈️ 👩🏿‍🚀 🧑🏿‍🚀 👨🏿‍🚀 👩🏿‍⚖️ 🧑🏿‍⚖️ 👨🏿‍⚖️ 👰🏿‍♀️ 👰🏿 👰🏿‍♂️ 🤵🏿‍♀️ 🤵🏿 🤵🏿‍♂️ 👸🏿 \U0001FAC5🏿 🤴🏿 🥷🏿 🦸🏿‍♀️ 🦸🏿 🦸🏿‍♂️ 🦹🏿‍♀️ 🦹🏿 🦹🏿‍♂️ 🤶🏿 🧑🏿‍🎄 🎅🏿 🧙🏿‍♀️ 🧙🏿 🧙🏿‍♂️ 🧝🏿‍♀️ 🧝🏿 🧝🏿‍♂️ 🧛🏿‍♀️ 🧛🏿 🧛🏿‍♂️ 🧜🏿‍♀️ 🧜🏿 🧜🏿‍♂️ 🧚🏿‍♀️ 🧚🏿 🧚🏿‍♂️ 👼🏿 🤰🏿 \U0001FAC4🏿 \U0001FAC3🏿 🤱🏿 👩🏿‍🍼 🧑🏿‍🍼 👨🏿‍🍼 🙇🏿‍♀️ 🙇🏿 🙇🏿‍♂️ 💁🏿‍♀️ 💁🏿 💁🏿‍♂️ 🙅🏿‍♀️ 🙅🏿 🙅🏿‍♂️ 🙆🏿‍♀️ 🙆🏿 🙆🏿‍♂️ 🙋🏿‍♀️ 🙋🏿 🙋🏿‍♂️ 🧏🏿‍♀️ 🧏🏿 🧏🏿‍♂️ 🤦🏿‍♀️ 🤦🏿 🤦🏿‍♂️ 🤷🏿‍♀️ 🤷🏿 🤷🏿‍♂️ 🙎🏿‍♀️ 🙎🏿 🙎🏿‍♂️ 🙍🏿‍♀️ 🙍🏿 🙍🏿‍♂️ 💇🏿‍♀️ 💇🏿 💇🏿‍♂️ 💆🏿‍♀️ 💆🏿 💆🏿‍♂️ 🧖🏿‍♀️ 🧖🏿 🧖🏿‍♂️ 💃🏿 🕺🏿 🕴🏿 👩🏿‍🦽 🧑🏿‍🦽 👨🏿‍🦽 👩🏿‍🦼 🧑🏿‍🦼 👨🏿‍🦼 🚶🏿‍♀️ 🚶🏿 🚶🏿‍♂️ 👩🏿‍🦯 🧑🏿‍🦯 👨🏿‍🦯 🧎🏿‍♀️ 🧎🏿 🧎🏿‍♂️ 🏃🏿‍♀️ 🏃🏿 🏃🏿‍♂️ 🧍🏿‍♀️ 🧍🏿 🧍🏿‍♂️ 👭🏿 🧑🏿‍🤝‍🧑🏿 👬🏿 👫🏿 🧗🏿‍♀️ 🧗🏿 🧗🏿‍♂️ 🏇🏿 🏂🏿 🏌🏿‍♀️ 🏌🏿 🏌🏿‍♂️ 🏄🏿‍♀️ 🏄🏿 🏄🏿‍♂️ 🚣🏿‍♀️ 🚣🏿 🚣🏿‍♂️ 🏊🏿‍♀️ 🏊🏿 🏊🏿‍♂️ ⛹🏿‍♀️ ⛹🏿 ⛹🏿‍♂️ 🏋🏿‍♀️ 🏋🏿 🏋🏿‍♂️ 🚴🏿‍♀️ 🚴🏿 🚴🏿‍♂️ 🚵🏿‍♀️ 🚵🏿 🚵🏿‍♂️ 🤸🏿‍♀️ 🤸🏿 🤸🏿‍♂️ 🤽🏿‍♀️ 🤽🏿 🤽🏿‍♂️ 🤾🏿‍♀️ 🤾🏿 🤾🏿‍♂️ 🤹🏿‍♀️ 🤹🏿 🤹🏿‍♂️ 🧘🏿‍♀️ 🧘🏿 🧘🏿‍♂️ 🛀🏿 🛌🏿 🐶 🐱 🐭 🐹 🐰 🦊 🐻 🐼 🐻‍❄️ 🐨 🐯 🦁 🐮 🐷 🐽 🐸 🐵 🙈 🙉 🙊 🐒 🐔 🐧 🐦 🐤 🐣 🐥 🦆 🦅 🦉 🦇 🐺 🐗 🐴 🦄 🐝 🪱 🐛 🦋 🐌 🐞 🐜 🪰 🪲 🪳 🦟 🦗 🕷 🕸 🦂 🐢 🐍 🦎 🦖 🦕 🐙 🦑 🦐 🦞 🦀 \U0001FAB8 🐡 🐠 🐟 🐬 🐳 🐋 🦈 🐊 🐅 🐆 🦓 🦍 🦧 🦣 🐘 🦛 🦏 🐪 🐫 🦒 🦘 🦬 🐃 🐂 🐄 🐎 🐖 🐏 🐑 🦙 🐐 🦌 🐕 🐩 🦮 🐕‍🦺 🐈 🐈‍⬛ 🪶 🐓 🦃 🦤 🦚 🦜 🦢 🦩 🕊 🐇 🦝 🦨 🦡 🦫 🦦 🦥 🐁 🐀 🐿 🦔 🐾 🐉 🐲 🌵 🎄 🌲 🌳 🌴 \U0001FAB9 \U0001FABA 🪵 🌱 🌿 ☘️ 🍀 🎍 🪴 🎋 🍃 🍂 🍁 🍄 🐚 🪨 🌾 💐 🌷 \U0001FAB7 🌹 🥀 🌺 🌸 🌼 🌻 🌞 🌝 🌛 🌜 🌚 🌕 🌖 🌗 🌘 🌑 🌒 🌓 🌔 🌙 🌎 🌍 🌏 🪐 💫 ⭐️ 🌟 ✨ ⚡️ ☄️ 💥 🔥 🌪 🌈 ☀️ 🌤 ⛅️ 🌥 ☁️ 🌦 🌧 🌩 🌨 ❄️ ☃️ ⛄️ 🌬 💨 💧 💦 \U0001FAE7 ☔️ ☂️ 🌊 🌫 🍏 🍎 🍐 🍊 🍋 🍌 🍉 🍇 🍓 🫐 🍈 🍒 🍑 🥭 🍍 🥥 🥝 🍅 🍆 🥑 🥦 🥬 🥒 🌶 🫑 🌽 🥕 🫒 🧄 🧅 🥔 🍠 \U0001FAD8 🥐 🥯 🍞 🥖 🥨 🧀 🥚 🍳 🧈 🥞 🧇 🥓 🥩 🍗 🍖 🦴 🌭 🍔 🍟 🍕 🫓 🥪 🥙 🧆 🌮 🌯 🫔 🥗 🥘 🫕 🥫 🍝 🍜 🍲 🍛 🍣 🍱 🥟 🦪 🍤 🍙 🍚 🍘 🍥 🥠 🥮 🍢 🍡 🍧 🍨 🍦 🥧 🧁 🍰 🎂 🍮 🍭 🍬 🍫 🍿 🍩 🍪 🌰 🥜 🍯 🥛 🍼 🫖 ☕️ 🍵 🧃 🥤 🧋 \U0001FAD9 🍶 🍺 🍻 🥂 🍷 \U0001FAD7 🥃 🍸 🍹 🧉 🍾 🧊 🥄 🍴 🍽 🥣 🥡 🥢 🧂 ⚽️ 🏀 🏈 ⚾️ 🥎 🎾 🏐 🏉 🥏 🎱 🪀 🏓 🏸 🏒 🏑 🥍 🏏 🪃 🥅 ⛳️ 🪁 🏹 🎣 🤿 🥊 🥋 🎽 🛹 🛼 🛷 ⛸ 🥌 🎿 ⛷ 🏂 🪂 🏋️‍♀️ 🏋️ 🏋️‍♂️ 🤼‍♀️ 🤼 🤼‍♂️ 🤸‍♀️ 🤸 🤸‍♂️ ⛹️‍♀️ ⛹️ ⛹️‍♂️ 🤺 🤾‍♀️ 🤾 🤾‍♂️ 🏌️‍♀️ 🏌️ 🏌️‍♂️ 🏇 🧘‍♀️ 🧘 🧘‍♂️ 🏄‍♀️ 🏄 🏄‍♂️ 🏊‍♀️ 🏊 🏊‍♂️ 🤽‍♀️ 🤽 🤽‍♂️ 🚣‍♀️ 🚣 🚣‍♂️ 🧗‍♀️ 🧗 🧗‍♂️ 🚵‍♀️ 🚵 🚵‍♂️ 🚴‍♀️ 🚴 🚴‍♂️ 🏆 🥇 🥈 🥉 🏅 🎖 🏵 🎗 🎫 🎟 🎪 🤹 🤹‍♂️ 🤹‍♀️ 🎭 🩰 🎨 🎬 🎤 🎧 🎼 🎹 🥁 🪘 🎷 🎺 🪗 🎸 🪕 🎻 🎲 ♟ 🎯 🎳 🎮 🎰 🧩 🚗 🚕 🚙 🚌 🚎 🏎 🚓 🚑 🚒 🚐 🛻 🚚 🚛 🚜 🦯 🦽 🦼 🛴 🚲 🛵 🏍 🛺 🚨 🚔 🚍 🚘 🚖 \U0001F6DE 🚡 🚠 🚟 🚃 🚋 🚞 🚝 🚄 🚅 🚈 🚂 🚆 🚇 🚊 🚉 ✈️ 🛫 🛬 🛩 💺 🛰 🚀 🛸 🚁 🛶 ⛵️ 🚤 🛥 🛳 ⛴ 🚢 ⚓️ \U0001F6DF 🪝 ⛽️ 🚧 🚦 🚥 🚏 🗺 🗿 🗽 🗼 🏰 🏯 🏟 🎡 🎢 \U0001F6DD 🎠 ⛲️ 🏖 🏝 🏜 🌋 ⛰ 🏔 🗻 🏕 ⛺️ 🛖 🏠 🏡 🏘 🏚 🏗 🏭 🏢 🏬 🏣 🏤 🏥 🏦 🏨 🏪 🏫 🏩 💒 🏛 ⛪️ 🕌 🕍 🛕 🕋 ⛩ 🛤 🛣 🗾 🎑 🏞 🌅 🌄 🌠 🎇 🎆 🌇 🌆 🏙 🌃 🌌 🌉 🌁 ⌚️ 📱 📲 💻 ⌨️ 🖥 🖨 🖱 🖲 🕹 🗜 💽 💾 💿 📀 📼 📷 📸 📹 🎥 📽 🎞 📞 ☎️ 📟 📠 📺 📻 🎙 🎚 🎛 🧭 ⏰ 🕰 ⌛️ ⏳ 📡 🔋 \U0001FAAB 🔌 💡 🔦 🕯 🪔 🧯 🛢 💸 💵 💴 💶 💷 🪙 💰 💳 💎 ⚖️ 🪜 🧰 🪛 🔧 🔨 ⚒ 🛠 ⛏ 🪚 🔩 ⚙️ 🪤 🧱 ⛓ 🧲 🔫 💣 🧨 🪓 🔪 🗡 ⚔️ 🛡 🚬 ⚰️ 🪦 ⚱️ 🏺 🔮 📿 🧿 \U0001FAAC 💈 ⚗️ 🔭 🔬 🕳 🩹 🩺 \U0001FA7B \U0001FA7C 💊 💉 🩸 🧬 🦠 🧫 🧪 🌡 🧹 🪠 🧺 🧻 🚽 🚰 🚿 🛁 🛀 🧼 🪥 🪒 🧽 🪣 🧴 🛎 🔑 🗝 🚪 🪑 🛋 🛏 🛌 🧸 🪆 🖼 🪞 🪟 🛍 🛒 🎁 🎈 🎏 🎀 🪄 🪅 🎊 🎉 \U0001FAA9 🎎 🏮 🎐 🧧 ✉️ 📩 📨 📧 💌 📥 📤 📦 🏷 🪧 📪 📫 📬 📭 📮 📯 📜 📃 📄 📑 🧾 📊 📈 📉 🗒 🗓 📆 📅 🗑 \U0001FAAA 📇 🗃 🗳 🗄 📋 📁 📂 🗂 🗞 📰 📓 📔 📒 📕 📗 📘 📙 📚 📖 🔖 🧷 🔗 📎 🖇 📐 📏 🧮 📌 📍 ✂️ 🖊 🖋 ✒️ 🖌 🖍 📝 ✏️ 🔍 🔎 🔏 🔐 🔒 🔓 ❤️ 🧡 💛 💚 💙 💜 🖤 🤍 🤎 ❤️‍🔥 ❤️‍🩹 💔 ❣️ 💕 💞 💓 💗 💖 💘 💝 💟 ☮️ ✝️ ☪️ 🕉 ☸️ ✡️ 🔯 🕎 ☯️ ☦️ 🛐 ⛎ ♈️ ♉️ ♊️ ♋️ ♌️ ♍️ ♎️ ♏️ ♐️ ♑️ ♒️ ♓️ 🆔 ⚛️ 🉑 ☢️ ☣️ 📴 📳 🈶 🈚️ 🈸 🈺 🈷️ ✴️ 🆚 💮 🉐 ㊙️ ㊗️ 🈴 🈵 🈹 🈲 🅰️ 🅱️ 🆎 🆑 🅾️ 🆘 ❌ ⭕️ 🛑 ⛔️ 📛 🚫 💯 💢 ♨️ 🚷 🚯 🚳 🚱 🔞 📵 🚭 ❗️ ❕ ❓ ❔ 🔅 🔆 〽️ ⚠️ 🚸 🔱 ⚜️ 🔰 ♻️ ✅ 🈯️ 💹 ❇️ ✳️ ❎ 🌐 💠 Ⓜ️ 🌀 💤 🏧 🚾 ♿️ 🅿️ 🛗 🈳 🈂️ 🛂 🛃 🛄 🛅 🚹 🚺 🚼 🚻 🚮 🎦 📶 🈁 🔣 ℹ️ 🔤 🔡 🔠 🆖 🆗 🆙 🆒 🆕 🆓 0️⃣ 1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣ 8️⃣ 9️⃣ 🔟 🔢 #️⃣ *️⃣ ⏏️ ▶️ ⏸ ⏯ ⏹ ⏺ ⏭ ⏮ ⏩ ⏪ ⏫ ⏬ ◀️ 🔼 🔽 ➡️ ⬅️ ⬆️ ⬇️ ↗️ ↘️ ↙️ ↖️ ↕️ ↔️ ↪️ ↩️ ⤴️ ⤵️ 🔀 🔁 🔂 🔄 🔃 🎵 🎶 ➕ ➖ ➗ ✖️ \U0001F7F0 ♾ 💲 💱 ™️ ©️ ®️ 〰️ ➰ ➿ 🔚 🔙 🔛 🔝 🔜 ✔️ 🔈 🔇 🔉 🔊 🔔 🔕 📣 📢 👁‍🗨 💬 💭 🗯 ♠️ ♣️ ♥️ ♦️ 🃏 🎴 🀄️ 🕐 🕑 🕒 🕓 🕔 🕕 🕖 🕗 🕘 🕙 🕚 🕛 🕜 🕝 🕞 🕟 🕠 🕡 🕢 🕣 🕤 🕥 🕦 🕧 "

var emojiList = strings.Split(emojis, " ")

var currentEmoji string

// You can set a custom seed here. The program will shuffle the emojiList based on this list. This means that, if the
// program restarts, the random list will be the same.
var seed = int64(000000000)

func setEmoji(num int) error {
	if num > len(emojiList) {
		return errors.New("num is longer than emoji list")
	}
	currentEmoji = emojiList[num]
	return nil
}

func GetCurrentEmoji() *string {
	return &currentEmoji
}

func ShuffleEmojiList() {
	// Before setting the seed, the current year is added to the seed. This means that if there is a new year, the
	// seed will change. This is needed because the emoji will be taken out of the list based on the day of the year.
	// When nothing changed about the seed after a year, the results will be exactly the same as the year before.
	random := rand.New(rand.NewSource(seed + int64(time.Now().Year())))

	random.Shuffle(len(emojiList), func(i, j int) {
		emojiList[i], emojiList[j] = emojiList[j], emojiList[i]
	})

	log.Println("Emoji list shuffled!")
}

func StartEmojiGenerator() {
	currentDay := time.Now()
	if currentDay.Day() == 1 && currentDay.Month() == 1 {
		ShuffleEmojiList()
	}

	err := setEmoji(currentDay.Day())
	if err != nil {
		log.Println("Error while setting emoji:", err)
	}
	log.Println("Emoji changed to " + currentEmoji + "!")

	nextDay := currentDay.AddDate(0, 0, 1)
	nextDayDate := time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, time.Local)

	diff := time.Duration(nextDay.Unix() - nextDayDate.Unix())

	time.Sleep(diff * time.Second)
	StartEmojiGenerator()
}