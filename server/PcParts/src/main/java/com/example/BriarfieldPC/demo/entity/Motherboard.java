package com.example.BriarfieldPC.demo.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.validation.constraints.NotBlank;

import static javax.persistence.GenerationType.IDENTITY;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
public class Motherboard {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    private long id;
    private String name;
    private String Manufacturer;
    private Integer price;
    private Integer ramSlots;
    private Integer driveSlots;
    private Integer maxRam;
    private String socket;
    private Integer TDP;
}
