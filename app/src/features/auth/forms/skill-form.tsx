import { FormField, FormItem, FormLabel, FormDescription, FormMessage, FormControl } from "@/components/ui/form"
import type { Control } from "react-hook-form"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { Label } from "@/components/ui/label"
import type { SkillFormValues } from "./types";

const skills: SkillLevelProps[] = [
  {
    value: "beginner",
    name: "Beginner",
    description: "You can boil water and heat up canned soup — and that's totally fine.",
  },
  {
    value: "intermediate",
    name: "Intermediate",
    description: "You're comfortable following recipes and can whip up simple meals, often using canned or packaged ingredients.",
  },
  {
    value: "advanced",
    name: "Advanced",
    description: "You cook from scratch often and feel confident experimenting with techniques and flavors.",
  },
  {
    value: "chef",
    name: "Chef",
    description: "You're a serious home cook or even work in the industry — nothing in the kitchen intimidates you.",
  },
];

export interface SkillFormProps {
    control: Control<SkillFormValues>
    showTitle?: boolean
}

export default function SkillForm({ control, showTitle = true }: SkillFormProps) {
    return (
        <FormField
            control={control}
            name="skill"
            render={({ field }) => (
                <FormItem>
                    {showTitle && <FormLabel>Current Skill Level</FormLabel>}
                    <FormDescription className="text-left">
                        How experienced are you in the kitchen? Don't worry, no matter your current skill level,
                        we'll help you level up.
                    </FormDescription>
                    <FormControl>
                        <RadioGroup
                            value={field.value}
                            onValueChange={field.onChange}
                        >
                            {skills.map(skill => (
                                <FormItem key={skill.value}>
                                    <FormControl>
                                        <SkillLevel {...skill} />
                                    </FormControl>
                                </FormItem>
                            ))}
                        </RadioGroup>
                    </FormControl>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}

interface SkillLevelProps {
    value: string
    name: string
    description: string
}

function SkillLevel({ value, name, description }: SkillLevelProps) {
    return (
        <div className="flex space-x-2 border rounded p-2">
            <RadioGroupItem value={value} id={value}/>
            <Label htmlFor={value}>
                <div className="flex flex-col items-start">
                    <span className="pb-2">{name}:</span>
                    <span className="text-left text-xs text-muted-foreground">
                        {description}
                    </span>
                </div>
            </Label>
        </div>
    )
}