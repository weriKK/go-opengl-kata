#version 410 core

in vec3 ourColor;
in vec2 TexCoord;

out vec4 color;

uniform sampler2D ourTexture0;
uniform sampler2D ourTexture1;

void main()
{
    vec4 tex0 = texture2D(ourTexture0, TexCoord);
    vec4 tex1 = texture2D(ourTexture1, TexCoord);

    color = mix(tex0, tex1, tex1.a);
}