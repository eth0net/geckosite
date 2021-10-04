"use strict";

class Gallery extends React.Component {
    constructor(props) {
        super(props);
        this.state = {index: 0};
        this.step = this.step.bind(this);
    }

    render() {
        const images = [
            ...this.props.images.slice(this.state.index),
            ...this.props.images.slice(0, this.state.index),
        ];
        const output = images.map(i => e("img", {key: i.src, src: i.src}));
        if (images.length > 1) output.push(
            e("button", {
                className: "prev",
                key: "prev",
                onClick: this.step.bind(this, -1),
            }, "❮"),
            e("button", {
                className: "next",
                key: "next",
                onClick: this.step.bind(this, 1),
            }, "❯")
        );
        return output;
    }

    step(increment, ev) {
        ev.preventDefault();
        const length = this.props.images.length;
        this.setState(s => {
            let index = s.index + increment;
            if (index < 0) {
                index += length;
            } else if (index >= length) {
                index -= length;
            }
            return {index};
        });
    }
}
